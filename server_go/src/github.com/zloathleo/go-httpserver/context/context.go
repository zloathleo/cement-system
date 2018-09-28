package context

import (
	"github.com/gin-gonic/gin"
	"github.com/zloathleo/go-httpserver/das"
)

func Init(router *gin.Engine) {
	router.Use(newCrosMiddleware())

	//pubGroup := router.Group("/")
	dasGroup := router.Group("/")
	//myGroup := router.Group("/my", newAuthMiddleware())

	//my.InitDashboard(pubGroup, dasGroup, myGroup)

	das.InitDashController(dasGroup)
}
