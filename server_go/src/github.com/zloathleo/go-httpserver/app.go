package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/zloathleo/go-httpserver/context"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("server is starting...")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Static("/public", "./public")
	context.Init(router)

	log.Println("server start at port 8088.")
	s := &http.Server{
		Addr:           ":8088",
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
