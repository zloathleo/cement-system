package my

import (
	"github.com/gin-gonic/gin"
	"github.com/zloathleo/go-httpserver/tiff"
	"net/http"
	"strconv"
)

func InitDashboard(pubGroup *gin.RouterGroup, dasGroup *gin.RouterGroup, myGroup *gin.RouterGroup) {

	//请求一张
	pubGroup.GET("/heatchart", func(c *gin.Context) {
		_indexString := c.DefaultQuery("index", "0")
		_pixelString := c.DefaultQuery("pixel", "8")
		_calc := c.DefaultQuery("calc", "w") //p w m

		_index, _ := strconv.Atoi(_indexString)
		_pixel, _ := strconv.Atoi(_pixelString)
		values, w, h := tiff.WalkTiffOne(_index, _pixel, _calc)

		c.JSON(http.StatusOK, gin.H{
			"w":    w,
			"h":    h,
			"data": values})
	})

	pubGroup.GET("/probabilitychart", func(c *gin.Context) {
		_periodString := c.DefaultQuery("period", "60")
		_pixelString := c.DefaultQuery("pixel", "8")
		_calc := c.DefaultQuery("calc", "w") //p w m

		_period, _ := strconv.Atoi(_periodString)
		_pixel, _ := strconv.Atoi(_pixelString)
		values, w, h := tiff.WalkTiffPeriod(_period, _pixel, _calc)

		c.JSON(http.StatusOK, gin.H{"w": w,
			"h":    h,
			"data": values})
	})

	pubGroup.GET("/trendchart", func(c *gin.Context) {
		c.JSON(http.StatusOK, []int{100, 200, 300, 400})
	})
}
