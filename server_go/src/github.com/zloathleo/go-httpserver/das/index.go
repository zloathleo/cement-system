package das

import (
	"fmt"
	"github.com/zloathleo/go-httpserver/common/config"
	"github.com/zloathleo/go-httpserver/common/logger"
)

func Init(){
	initStore()
	dasUrl = fmt.Sprintf("http://%v:%d",config.AppConfig.Das.Host,config.AppConfig.Das.Port)
	logger.Printf("das url %v.",dasUrl)
}

