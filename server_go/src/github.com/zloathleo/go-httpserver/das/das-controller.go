// http接口

package das

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type PushDas struct {
	Rows      []Das `json:"rows"`
	TimeStamp int64 `json:"timestamp"`
}

type Das struct {
	PointName string  `json:"name"`
	Value     float64 `json:"value"`
	TimeStamp int64   `json:"timestamp"`
}

func InitDashController(dasGroup *gin.RouterGroup) {

	//客户端请求实时数据
	dasGroup.GET("/data", func(c *gin.Context) {
		//dashboard,control
		_page := c.DefaultQuery("page", "control")
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"page": _page,
		})
		return
	})

	//客户端请求控制指令
	dasGroup.POST("/control", func(c *gin.Context) {
		dataJson := c.PostForm("content")
		log.Printf("control data %v", dataJson)
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
			log.Println(err)
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
		log.Printf("push data %v", dataJson)
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
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"code": -1,
				"msg":  fmt.Sprintf("param 'content' [%v] is err [%v]", dataJson, err),
			})
			return
		}
	})

}
