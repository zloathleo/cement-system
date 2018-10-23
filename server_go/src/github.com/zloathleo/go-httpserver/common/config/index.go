package config

import (
	"github.com/zloathleo/go-httpserver/common/logger"
)

const LogoAsic = `
***********************************************************************************
***********************************************************************************
****                             ****            ************                  ****
****                             ******* *********************                 ****
****                              ***************************                  ****
****                             ******************    ***                     ****
****                          ******************     ***                       ****
****                       ******************    *****                         ****
****           *****     **********    ************                            ****
****            ***************************    ***                             ****
****               ************      ****                                      ****
****                 ****            *                                         ****
****                 ***                                                       ****
****                ***                                                        ****
***********************************************************************************
***********************************************************************************
**** load config file ok
**** App.AutoLogic %v
**** Das.Host %s
**** Das.Port %d 
***********************************************************************************

`

const uiconfiglog = `
***********************************************************************************
**** load ui config file ok
**** Dashboard.LineChartPoints %s 
***********************************************************************************
`

var AppConfig = struct {
	App struct {
		Name         string `default:"center-server"`
		Port         uint   `required:"true" default:"8088"`
		AutoLogic    bool   `default:false`
		LogLevel     string `default:"W"`
	}
	TempvisionDas struct {
		Disable bool `default:false`
		Host string `default:"localhost"`
		Port uint   `default:"8080"`
	} `yaml:"tempvision-das"`

	PointValue []struct {
		Name       string `required:"true"`
		Conversion string `required:"true"`
	} `yaml:"point-value"`

	CmdParam []struct {
		Name       string `required:"true"`
		Conversion string `required:"true"`
	} `yaml:"cmd-param"`

	CmdTarget []struct {
		Name       string `required:"true"`
		Conversion string `required:"true"`
	} `yaml:"cmd-target"`
}{}

var UIConfig = struct {
	Dashboard struct {
		LineChartPoints string
		DefaultDur      int `default:"3600"`
	}
}{}

func Init() {
	err := Load(&AppConfig, "config.yml")
	if err != nil {
		logger.Fatalf("load config file error [%v]", err)
	} else {
		logger.Infof(LogoAsic, AppConfig.App.AutoLogic, AppConfig.TempvisionDas.Host, AppConfig.TempvisionDas.Port)
	}
	err = Load(&UIConfig, "ui.yml")
	if err != nil {
		logger.Fatalf("load ui config file error [%v]", err)
	} else {
		logger.Infof(uiconfiglog, UIConfig.Dashboard.LineChartPoints)
	}
	parsePointConfig()

}
