// 和底层接口通讯
package das

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/zloathleo/go-httpserver/common/logger"
	"net/http"
	"net/url"
)

func checkPostCommand(pushDas *PushDas) error {
	rows := pushDas.Rows
	if len(rows) == 1 {
		row := rows[0]
		return checkSingleCommand(row.PointName, row.Value)
	} else {
		return errors.New("command rows is not 1.")
	}
}

func checkSingleCommand(pn string, value float64) error {
	switch pn {
	case "2_00019":
		return check200019(value)
	case "2_00017":
		return check200017(value)
	case "2_00020":
		return check200020(value)
	case "2_00018":
		return check200018(value)
	default:
		return nil
	}
}

//探头A启动条件
func check200019(value float64) error {
	if value == 1 {
		//冷却风压力 0.1MPa
		err := validPoint("4_40002")
		if err != nil {
			return err
		}

		//A探头内部温度 55
		err = validPoint("5_40013")
		if err != nil {
			return err
		}
	}
	return nil
}

//探头A进退条件
func check200017(value float64) error {
	//是否就地模式
	err := validPoint("1_00001")
	if err != nil {
		return err
	}

	//是否远程模式
	err = validPoint("1_00002")
	if err != nil {
		return err
	}

	//前进
	if value == 1 {
		//吹扫风压力 5KPa
		err = validPoint("4_40001")
		if err != nil {
			return err
		}

		//冷却风压力
		err = validPoint("4_40002")
		if err != nil {
			return err
		}

		//A探头内部温度
		err = validPoint("5_40013")
		if err != nil {
			return err
		}
	}
	return nil
}

//探头B启动条件
func check200020(value float64) error {
	if value == 1 {
		//冷却风压力 0.1MPa
		err := validPoint("4_40002")
		if err != nil {
			return err
		}

		//A探头内部温度 55
		err = validPoint("5_40083")
		if err != nil {
			return err
		}
	}
	return nil
}

//探头B进退条件
func check200018(value float64) error {
	//是否就地模式
	err := validPoint("1_00001")
	if err != nil {
		return err
	}

	//是否远程模式
	err = validPoint("1_00002")
	if err != nil {
		return err
	}

	//前进
	if value == 1 {
		//吹扫风压力 5KPa
		err = validPoint("4_40001")
		if err != nil {
			return err
		}

		//冷却风压力
		err = validPoint("4_40002")
		if err != nil {
			return err
		}

		//A探头内部温度
		err = validPoint("5_40083")
		if err != nil {
			return err
		}
	}
	return nil
}

//命令下发到接口程序
func cmdPostDDI(dataJson string) (*http.Response, error) {
	//http request to ddi
	resp, err := http.PostForm(dasUrl+"/control",
		url.Values{"content": {dataJson}})
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		logger.Printf("命令下发异常,[%v].",err)
		return resp, err
	} else {
		return resp, nil
	}
}

//发送控制命令到底层接口
func sendData2DDI(dataJson string) (*http.Response, error) {
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
