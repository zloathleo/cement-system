//存储历史数据
package das

import (
	"github.com/coocood/freecache"
	"runtime/debug"
	"github.com/zloathleo/go-httpserver/utils"
	"github.com/zloathleo/go-httpserver/common/logger"
	"github.com/zloathleo/go-httpserver/common/config"
	"gopkg.in/Knetic/govaluate.v2"
)

var cache *freecache.Cache
var pointsConfigMap = make(map[string]*govaluate.EvaluableExpression)
func initStore(){
	//10M
	cacheSize := 10 * 1024 * 1024
	cache = freecache.NewCache(cacheSize)
	debug.SetGCPercent(10)

	//config
	points := config.AppConfig.Point
	for _, v := range points {
		conversion := v.Conversion
		if conversion != ""{
			expression, err := govaluate.NewEvaluableExpression(conversion)
			if err != nil {
				logger.Fatalf("%v conversion config $v is error.",v.Name,v.Conversion)
			}else{
				pointsConfigMap[v.Name] = expression
			}
		}
	}
	logger.Printf("point config load ok [%v].",pointsConfigMap)
}

func saveHistoryData(pushDas *PushDas) error {
	rows := pushDas.Rows
	timeStamp := pushDas.TimeStamp

	logger.Println("cache timeStamp data:",timeStamp)
	for _,item := range rows {
		saveValue(item)
	}
	return nil
}

func saveValue(item *Das){
	key := item.PointName
	value := item.Value
	conversion := pointsConfigMap[key]
	if conversion != nil {
		parameters := make(map[string]interface{}, 8)
		parameters["x"] = value

		result, err := conversion.Evaluate(parameters)
		if err == nil {
			value = result.(float64)
		}else{
			logger.Printf("conversion %v for %v is error",conversion.String(),value)
		}
	}
	cache.Set([]byte(key),utils.Float64ToBytes(value),0)
}

func getCurrentData(id string) (bool,float64){
	valueBytes,_ := cache.Get([]byte(id))
	if valueBytes != nil {
		return true,utils.Float64FromBytes(valueBytes)
	}else{
		return false,0
	}

}
