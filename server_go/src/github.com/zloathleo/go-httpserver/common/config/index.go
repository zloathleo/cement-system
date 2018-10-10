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
**** Das.Host %v
**** Das.Port %v 
***********************************************************************************
`

var AppConfig = struct {

	App struct {
		Name     string  `default:"center-server"`
		Port     uint   `required:"true" default:"8088"`
		//Password string `required:"true" env:"DBPassword"`
	}

	Das struct {
		Host     string `default:"localhost"`
		Port     uint   `default:"8088"`
	}

	Point []struct {
		Name  string `required:"true"`
		Conversion string `required:"true"`
	}
}{}

func Init(){
		err := Load(&AppConfig, "config.yml")
		if err != nil{
			logger.Fatalf("load config file error [%v]",err)
		} else{
			logger.Printf(LogoAsic,AppConfig.Das.Host,AppConfig.Das.Port)
		}
}

