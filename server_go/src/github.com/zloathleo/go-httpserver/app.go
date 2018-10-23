package main

import (
	"github.com/zloathleo/go-httpserver/common/logger"

	"github.com/zloathleo/go-httpserver/common"
	"github.com/zloathleo/go-httpserver/common/config"
	"github.com/zloathleo/go-httpserver/httpr"

	"os"
	"os/signal"

	"github.com/zloathleo/go-httpserver/history2"
)

func main() {
	common.Init()
	logger.Warnf("%v is starting ...", config.AppConfig.App.Name)
	history2.Init()
	httpr.InitWWW()
	//common.NotifyExit()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	cs := <-c
	history2.Exit()
	logger.Infof("Got signal [%v], app exit.", cs)
}
