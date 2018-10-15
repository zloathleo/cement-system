//存储历史数据
package das

import (
	"github.com/coocood/freecache"
	"github.com/gin-gonic/gin/json"
	"github.com/zloathleo/go-httpserver/common/config"
	"github.com/zloathleo/go-httpserver/common/logger"
	"github.com/zloathleo/go-httpserver/utils"
	"gopkg.in/Knetic/govaluate.v2"
	"math"
	"runtime/debug"
)

var cache *freecache.Cache
var pointsConfigMap = make(map[string]*govaluate.EvaluableExpression)

func initStore() {
	//10M
	cacheSize := 10 * 1024 * 1024
	cache = freecache.NewCache(cacheSize)
	debug.SetGCPercent(10)

	//config
	points := config.AppConfig.Point
	for _, v := range points {
		conversion := v.Conversion
		if conversion != "" {
			expression, err := govaluate.NewEvaluableExpression(conversion)
			if err != nil {
				logger.Fatalf("%v conversion config $v is error.", v.Name, v.Conversion)
			} else {
				pointsConfigMap[v.Name] = expression
			}
		}
	}
	//logger.Printf("point config load ok [%v].", pointsConfigMap)
}

func saveHistoryData(pushDas *PushDas) error {
	rows := pushDas.Rows
	timeStamp := pushDas.TimeStamp

	logger.Println("cache timeStamp data:", timeStamp)
	for _, item := range rows {
		saveValue(item)
	}

	//自动退出
	if config.AppConfig.App.Autologic {
		//检查环境退出探头或关闭电源
		autoBackCheckStatus()
	}

	return nil
}

//checkCount > 10 次后 才生效
var checkCount = 0

//自动退出 检查环境退出探头或关闭电源
func autoBackCheckStatus() {
	if checkCount < 20 {
		checkCount++
		return
	} else  {
		if checkCount < 1000 {
			checkCount++
		}
	}

	//A探头退到位
	_, at := getCurrentData("1_00005")
	////B探头退到位
	//_, bt := getCurrentData("1_00007")

	//吹扫风合法值>=5
	e1 := validPoint("4_40001")
	//冷却风合法
	e2 := validPoint("4_40002")

	//镜头A温度<=55
	e3 := validPoint("4_40003")
	//探头A内部温度<=55
	e4 := validPoint("5_40013")

	////镜头B温度<=55
	//e5 := validPoint("4_40004")
	////探头B内部温度<=55
	//e6 := validPoint("5_40083")

	var rows []*Das

	logger.Println("auto logic start------------------------")
	if e3 != nil || e4 != nil {
		//高温
		rows = append(rows, &Das{"3_00019", 0, 1})
		logger.Printf("A high alarm ,reason is [%v,%v].", e3, e4)
	}
	if e2 != nil || e4 != nil {
		//A探头电源关闭
		rows = append(rows, &Das{"2_00019", 0, 0})
		logger.Printf("A Cut down,reason is [%v,%v].", e2, e4)
	} else {
		//A探头正常
		rows = append(rows, &Das{"3_00017", 0, 1})
		logger.Println("A ok")
	}
	if (e1 != nil || e2 != nil || e3 != nil || e4 != nil) && at != 1 {
		//A镜头退出
		rows = append(rows, &Das{"2_00017", 0, 0})
		logger.Printf("A return back,reason is[%v,%v,%v,%v].", e1, e2, e3, e4)
	}
	logger.Println("auto logic end------------------------")

	//b探头暂停
	//if e2 != nil || e6 != nil {
	//	//B探头电源关闭
	//	rows = append(rows, &Das{"2_00020", 0, 0})
	//}
	//if e1 != nil || e2 != nil || e5 != nil || bt!= 1{
	//	//B镜头退出
	//	rows = append(rows, &Das{"2_00018", 0, 0})
	//}

	if len(rows) > 0 {
		pushDas := PushDas{rows, 0}
		pushDasBytes, err := json.Marshal(pushDas)
		if err != nil {
			logger.Printf("saveHistoryData checkStatus error %v", err)
		} else {
			cmdPostDDI(string(pushDasBytes))
		}
	}

}

func saveValue(item *Das) {
	key := item.PointName
	value := item.Value
	conversion := pointsConfigMap[key]
	if conversion != nil {
		parameters := make(map[string]interface{}, 8)
		parameters["x"] = value

		result, err := conversion.Evaluate(parameters)
		if err == nil {
			value = result.(float64)
		} else {
			logger.Printf("conversion %v for %v is error", conversion.String(), value)
		}
	}
	if math.Abs(value-0) < 0.00001 {
		value = 0
	}
	//60s 超时
	cache.Set([]byte(key), utils.Float64ToBytes(value), 600)
}

func getCurrentData(id string) (bool, float64) {
	valueBytes, _ := cache.Get([]byte(id))
	if valueBytes != nil {
		return true, utils.Float64FromBytes(valueBytes)
	} else {
		return false, 0
	}
}
