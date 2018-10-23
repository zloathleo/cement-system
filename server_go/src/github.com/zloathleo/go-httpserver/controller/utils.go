package controller

import (
	"github.com/zloathleo/go-httpserver/common/config"
	"fmt"
	"github.com/pkg/errors"
)

//验证测点值合法性
func validPointValue(key string, value float64) error {
	conversion := config.CmdParamMap[key]
	if conversion != nil {
		parameters := make(map[string]interface{}, 4)
		parameters["x"] = value
		result, err := conversion.Evaluate(parameters)
		if err == nil {
			r := result.(bool)
			if r {
				return nil
			} else {
				return errors.New(fmt.Sprintf("point '%v' conversion %v for %v is valid.", key, conversion.String(), value))
			}
		} else {
			return errors.New(fmt.Sprintf("conversion %v for %v is error", conversion.String(), value))
		}
	} else {
		return nil
	}
}
