package das

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/zloathleo/go-httpserver/common/config"
	"github.com/zloathleo/go-httpserver/common/logger"
	"gopkg.in/Knetic/govaluate.v2"
)

//验证测点是否正常
var cmdValidMap = make(map[string]*govaluate.EvaluableExpression)

//命令执行组态条件配置
func initCmdValid() {
	//config
	points := config.AppConfig.Cmd
	for _, v := range points {
		conversion := v.Conversion
		if conversion != "" {
			expression, err := govaluate.NewEvaluableExpression(conversion)
			if err != nil {
				logger.Fatalf("%v cmd valid config $v is error.", v.Name, v.Conversion)
			} else {
				cmdValidMap[v.Name] = expression
			}
		}
	}
	//logger.Printf("cmd valid load ok [%v].", cmdValidMap)
}

func validPoint(key string) error {
	exist, value := getCurrentData(key)
	if !exist {
		return errors.New(fmt.Sprintf("point '%v' value is not exist.", key))
	}
	conversion := cmdValidMap[key]
	if conversion != nil {
		parameters := make(map[string]interface{}, 8)
		parameters["x"] = value

		result, err := conversion.Evaluate(parameters)
		if err == nil {
			r := result.(bool)
			if r {
				return nil
			} else {
				return errors.New(fmt.Sprintf("point '%v' conversion %v for %v is valid.", key, conversion.String() ,value))
			}
		} else {
			return errors.New(fmt.Sprintf("conversion %v for %v is error", conversion.String(), value))
		}
	} else {
		return nil
	}
}

//
////吹扫风压力 低压 5KPa
//func valid440001() error {
//	return validPoint("4_40001")
//}
//
////冷却风压力 0.1MPa  0.35MPa
//func valid440002() error {
//	exist1, value1 := getCurrentData("4_40002")
//	if !exist1 {
//		return errors.New("point '4_40002' value is not exist.")
//	}
//	if value1 < 0.1 {
//		return errors.New("point '4_40002' value < 0.1.")
//	}
//	if value1 > 0.35 {
//		return errors.New("point '4_40002' value < 0.1.")
//	}
//	return nil
//}
//
////A镜头温度 高温 55
//func valid540013() error {
//	exist, value := getCurrentData("4_40003")
//	if !exist {
//		return errors.New("point '5_40013' value is not exist.")
//	}
//	if value > 55 {
//		return errors.New("point '5_40013' value > 55.")
//	}
//	return nil
//}
//
////A探头内部温度 高温 55
//func valid540013() error {
//	exist, value := getCurrentData("5_40013")
//	if !exist {
//		return errors.New("point '5_40013' value is not exist.")
//	}
//	if value > 55 {
//		return errors.New("point '5_40013' value > 55.")
//	}
//	return nil
//}
//
////是否就地模式
//func valid100001() error {
//	exist1, value1 := getCurrentData("1_00001")
//	if !exist1 {
//		return errors.New("point '1_00001' value is not exist.")
//	}
//	if value1 == 1 {
//		return errors.New("point '1_00001' value = 1.")
//	}
//	return nil
//}
//
////是否远程模式
//func valid100002() error {
//	exist, value := getCurrentData("1_00002")
//	if !exist {
//		return errors.New("point '1_00002' value is not exist.")
//	}
//	if value == 0 {
//		return errors.New("point '1_00002' value = 0.")
//	}
//	return nil
//}
//
///////////////////////////////////////////////
