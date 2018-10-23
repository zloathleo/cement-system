// http接口
package controller

import (
	"encoding/json"
	"fmt"

	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	. "github.com/zloathleo/go-httpserver/dstruct"

	"github.com/zloathleo/go-httpserver/common/logger"
)

func InitDashController(dasGroup *gin.RouterGroup) {
	InitCmdConfig()

	//调试客户端请求页面数据
	dasGroup.GET("/data", func(c *gin.Context) {
		pageData := serviceRenderPageControl()
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"rows": pageData.Rows,
		})
	})

	//客户端请求页面数据
	dasGroup.GET("/page", func(c *gin.Context) {
		//dashboard,control
		_page := c.DefaultQuery("name", "control")
		var pageData *PushDas
		if strings.EqualFold(_page, "dashboard") {
			pageData = serviceRenderPageDashboard()
		} else if strings.EqualFold(_page, "control") {
			pageData = serviceRenderPageControl()
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": -1,
				"msg":  fmt.Sprintf("param 'name' [%v] is err ", _page),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"rows": pageData.Rows,
		})
	})

	//客户端请求历史曲线
	dasGroup.GET("/his-chart", func(c *gin.Context) {
		//dashboard,control
		points := c.Query("points")
		to := c.Query("to")   //结束时间
		dur := c.Query("dur") //时长

		pointsArray := strings.Split(points, ",")
		if pointsArray == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": -1,
				"msg":  fmt.Sprintf("param 'points' [%s] is err.", points),
			})
			return
		}

		var to64 int64
		if to == "" {
			to64 = time.Now().Unix()
		} else {
			var _err error
			to64, _err = strconv.ParseInt(to, 10, 64)
			if _err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code": -1,
					"msg":  fmt.Sprintf("param 'to' [%s] is err.", to),
				})
				return
			}
		}

		durInt, err := strconv.Atoi(dur)
		if err != nil {
			durInt = 600
		}

		begin := time.Now().Unix()
		jsonBuffer := fetchHistoryChartData(pointsArray, to64, durInt)
		logger.Debugf("req his const: %d", int(time.Now().Unix()-begin))

		c.Status(http.StatusOK)
		c.Writer.Write(jsonBuffer.Bytes())

	})

	//客户端请求控制指令
	dasGroup.POST("/control", func(c *gin.Context) {
		dataJson := c.PostForm("content")
		logger.Infof("control data %v", dataJson)
		if len(dataJson) < 5 {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": -1,
				"msg":  fmt.Sprintf("param 'content' [%v] is err ", dataJson),
			})
			return
		}

		//发送到底层
		res, err := serviceSendCmd(dataJson)
		logger.Infof("control command response %v", res)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": -1,
				"msg":  fmt.Sprintf("cmd exception [%v]", err),
			})
		}
	})

	//下位机上传数据
	dasGroup.POST("/push", func(c *gin.Context) {
		dataJson := c.PostForm("content")
		//logger.Infof("push data %v.", dataJson)
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
			serviceSaveDas(&pushDas)
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
			})
			return
		} else {
			logger.Warnln(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"code": -1,
				"msg":  fmt.Sprintf("param 'content' [%v] is err [%v]", dataJson, err),
			})
			return
		}
	})

}
