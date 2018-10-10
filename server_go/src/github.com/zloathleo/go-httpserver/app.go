package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/zloathleo/go-httpserver/context"
	"net/http"
	"time"
	"fmt"
	"github.com/zloathleo/go-httpserver/common/logger"
	"github.com/zloathleo/go-httpserver/common/config"
	"github.com/zloathleo/go-httpserver/das"
)

func main() {
	logger.Init()
	config.Init()
	logger.Printf("%v is starting ...",config.AppConfig.App.Name)
	das.Init()

	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = logger.GetOutput()

	router := gin.New()
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Static("/public", "./public")
	context.Init(router)


	addr := fmt.Sprintf(":%d", config.AppConfig.App.Port)
	s := &http.Server{
		Addr:           addr ,
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
	logger.Printf("%v http listen at %d.",config.AppConfig.App.Name,config.AppConfig.App.Port)
}
