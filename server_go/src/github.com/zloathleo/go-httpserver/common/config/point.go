package config

import (
	"gopkg.in/Knetic/govaluate.v2"
	"github.com/zloathleo/go-httpserver/common/logger"
)

//测点获取值变换
var PointValueMap = make(map[string]*govaluate.EvaluableExpression)

//指令相关测点条件
var CmdParamMap = make(map[string]*govaluate.EvaluableExpression)

////指令目标逻辑
//var CmdTargetMap = make(map[string]*govaluate.EvaluableExpression)

//var cmdTargetfunctions map[string]govaluate.ExpressionFunction
//
//func init() {
//	cmdTargetfunctions = map[string]govaluate.ExpressionFunction{
//		"valid": func(args ...interface{}) (interface{}, error) {
//			key := args[0].(string)
//			logger.Println("key:", key)
//			return true, nil
//		},
//	}
//}

func parsePointConfig() {
	points := AppConfig.PointValue
	for _, v := range points {
		conversion := v.Conversion
		if conversion != "" {
			expression, err := govaluate.NewEvaluableExpression(conversion)
			if err != nil {
				logger.Fatalf("%v point value config %v is error.", v.Name, v.Conversion)
			} else {
				PointValueMap[v.Name] = expression
			}
		}
	}

	logger.Info("config point conversion ok.")

	cmdParam := AppConfig.CmdParam
	for _, v := range cmdParam {
		conversion := v.Conversion
		if conversion != "" {
			expression, err := govaluate.NewEvaluableExpression(conversion)
			if err != nil {
				logger.Fatalf("%v cmd param config %v is error.", v.Name, v.Conversion)
			} else {
				CmdParamMap[v.Name] = expression
			}
		}
	}
	logger.Info("config cmd param conversion ok.")

	//cmdTarget := AppConfig.CmdTarget
	//for _, v := range cmdTarget {
	//	conversion := v.Conversion
	//	if conversion != "" {
	//		expression, err := govaluate.NewEvaluableExpressionWithFunctions(conversion, cmdTargetfunctions)
	//		if err != nil {
	//			logger.Errorf("%v cmd target config %v is error.", v.Name, v.Conversion)
	//		} else {
	//			CmdTargetMap[v.Name] = expression
	//		}
	//	}
	//}
	//logger.Println("config cmd target conversion ok.")
}
