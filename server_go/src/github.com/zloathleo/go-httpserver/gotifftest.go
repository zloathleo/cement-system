package main

import (
	"time"
	"github.com/robfig/cron"
	"github.com/zloathleo/go-httpserver/common/logger"
	"os"
	"os/signal"
)

var cornSpecCreateTableInit = "0 * * * * ?" //每隔11分钟执行一次

func main(){
	//aaa()

	//history2.Init()

	//todayFileName := utils.GetDay24Time(time.Now())
	//fmt.Println(todayFileName)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	cs := <-c
	logger.Infof("Got signal [%v], app exit.", cs)
}


func aaa() {

	c := cron.New()

	c.AddFunc(cornSpecCreateTableInit, func() {
		logger.Println("cron running:", time.Now())

	})

	c.Start()
}

