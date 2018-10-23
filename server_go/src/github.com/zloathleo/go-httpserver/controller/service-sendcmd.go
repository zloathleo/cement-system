/**
控制命令逻辑
*/
package controller

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zloathleo/go-httpserver/common/config"
	. "github.com/zloathleo/go-httpserver/dstruct"
	"github.com/zloathleo/go-httpserver/memcache"
	"gopkg.in/Knetic/govaluate.v2"
	"github.com/zloathleo/go-httpserver/common/logger"
	"net/http"
)

var (
	//指令目标逻辑
	cmdTargetMap = make(map[string]*govaluate.EvaluableExpression)
	//eval绑定的方法
	cmdTargetfunctions = make(map[string]govaluate.ExpressionFunction)

	//指令下发接口程序请求url
	cmdDasTargetUrl = ""
)

func InitCmdConfig() {
	cmdDasTargetUrl = fmt.Sprintf("http://%v:%d", config.AppConfig.TempvisionDas.Host, config.AppConfig.TempvisionDas.Port)
	logger.Infof("Tempvision das url %v.", cmdDasTargetUrl)

	cmdTargetfunctions["valid"] = func(args ...interface{}) (interface{}, error) {
		key := args[0].(string)
		err := validPoint(key)
		if err != nil {
			return false, err
		} else {
			return true, nil
		}
	}

	cmdTarget := config.AppConfig.CmdTarget
	for _, v := range cmdTarget {
		conversion := v.Conversion
		if conversion != "" {
			expression, err := govaluate.NewEvaluableExpressionWithFunctions(conversion, cmdTargetfunctions)
			if err != nil {
				logger.Fatalf("%v cmd target config %v is error.", v.Name, v.Conversion)
			} else {
				cmdTargetMap[v.Name] = expression
			}
		}
	}
	logger.Info("config cmd target conversion ok.")
}

//发送控制命令到底层接口
func serviceSendCmd(dataJson string) (*http.Response, error) {
	var pushDas PushDas
	err := json.Unmarshal([]byte(dataJson), &pushDas)
	if err == nil {
		err = checkPostCommand(&pushDas)
		if err == nil {
			return cmdPostDDI(dataJson)
		}
	}
	return nil, err
}

func validPoint(key string) error {
	if key == "" {
		return errors.New("point key is nil.")
	} else {
		exist, value := memcache.GetCurrentData(key)
		if !exist {
			return errors.New(fmt.Sprintf("point '%v' value is not exist.", key))
		}
		return validPointValue(key, value)
	}
}

//检查下发命令
func checkPostCommand(pushDas *PushDas) error {
	rows := pushDas.Rows
	if len(rows) == 1 {
		row := rows[0]
		return checkSingleCommand(row.PointName, row.Value)
	} else {
		return errors.New("command rows is not 1.")
	}
}

//检查下发命令
func checkSingleCommand(pn string, value float64) error {
	conversion := cmdTargetMap[pn]
	if conversion != nil {
		parameters := make(map[string]interface{}, 8)
		parameters["x"] = value
		result, err := conversion.Evaluate(parameters)
		if err == nil {
			ret := result.(bool)
			if ret {
				return nil
			} else {
				return errors.New(fmt.Sprintf("conversion [%v] eva false", conversion.String()))
			}
		} else {
			return err
		}
	}
	return nil
}
