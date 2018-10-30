package common

import (
	"github.com/zloathleo/go-httpserver/common/config"
	"github.com/zloathleo/go-httpserver/common/logger"
	"math"
)

func Init() {
	logger.Init()
	config.Init()
	logger.SetLevel(config.AppConfig.App.LogLevel)
}

//值转换
func TranslatePointValue(key string, value float64) float64 {
	if config.AppConfig.App.IgnoreValueZero {
		if math.Abs(value-0) < 0.00001 {
			//如果设置忽略0值
			return math.NaN()
		}
	}

	conversion := config.PointValueMap[key]
	if conversion != nil {
		parameters := make(map[string]interface{}, 8)
		parameters["x"] = value
		result, err := conversion.Evaluate(parameters)
		if err == nil {
			value = result.(float64)
		} else {
			logger.Infof("conversion %v for %v is error", conversion.String(), value)
		}
	}
	if math.Abs(value-0) < 0.00001 {
		value = 0
	}
	return value
}
