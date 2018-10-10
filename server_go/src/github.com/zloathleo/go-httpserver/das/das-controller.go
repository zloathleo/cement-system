// http接口

package das

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"github.com/zloathleo/go-httpserver/common/logger"
)

func InitDashController(dasGroup *gin.RouterGroup) {

	//客户端请求页面数据
	dasGroup.GET("/page", func(c *gin.Context) {
		//dashboard,control
		_page := c.DefaultQuery("name", "control")
		var pageData *PushDas
		if strings.EqualFold(_page,"dashboard"){
			pageData = renderPageDashboard()
		}else if strings.EqualFold(_page,"control"){
			pageData = renderPageControl()
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"rows": pageData.Rows,
		})
		return
	})

	//客户端请求控制指令
	dasGroup.POST("/control", func(c *gin.Context) {
		dataJson := c.PostForm("content")
		logger.Printf("control data %v", dataJson)
		if len(dataJson) < 5 {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": -1,
				"msg":  fmt.Sprintf("param 'content' [%v] is err ", dataJson),
			})
			return
		}
		var pushDas PushDas
		err := json.Unmarshal([]byte(dataJson), &pushDas)
		if err == nil {
			//发送到底层
			sendData2DDI(dataJson)
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
			})
			return
		} else {
			logger.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"code": -1,
				"msg":  fmt.Sprintf("param 'content' [%v] is err [%v]", dataJson, err),
			})
			return
		}
	})

	//下位机上传数据
	dasGroup.POST("/push", func(c *gin.Context) {
		dataJson := c.PostForm("content")
		//logger.Printf("push data %v", dataJson)
		if len(dataJson) < 5 {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": -1,
				"msg":  fmt.Sprintf("param 'content' [%v] is err ", dataJson),
			})
			return
		}
		var pushDas PushDas
		err := json.Unmarshal([]byte(dataJson), &pushDas)
		if err == nil {
			//保存到历史数据库
			saveHistoryData(&pushDas)
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
			})
			return
		} else {
			logger.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"code": -1,
				"msg":  fmt.Sprintf("param 'content' [%v] is err [%v]", dataJson, err),
			})
			return
		}
	})

}
