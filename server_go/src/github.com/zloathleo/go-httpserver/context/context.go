package context

import (
	"github.com/gin-gonic/gin"
	"github.com/zloathleo/go-httpserver/das"
)

func Init(router *gin.Engine) {
	router.Use(newCrosMiddleware())
	dasGroup := router.Group("/")
	das.InitDashController(dasGroup)
}
