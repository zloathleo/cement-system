package das

import (
	"fmt"
	"github.com/zloathleo/go-httpserver/common/config"
	"github.com/zloathleo/go-httpserver/common/logger"
)

//底层接口通讯软件ip和端口
var dasUrl = ""
func Init(){
	initStore()
	initCmdValid()
	dasUrl = fmt.Sprintf("http://%v:%d",config.AppConfig.Das.Host,config.AppConfig.Das.Port)
	logger.Printf("das url %v.",dasUrl)
}

