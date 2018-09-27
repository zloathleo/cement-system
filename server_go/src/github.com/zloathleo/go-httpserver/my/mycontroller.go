// 个人账户管理

package my

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(pubGroup *gin.RouterGroup, myGroup *gin.RouterGroup) {

	//登录
	pubGroup.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.Bind(&json); err == nil {
			if json.User == "manu" && json.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"data": gin.H{
					"firstname": json.User,
					"lastname":  json.Password,
				}})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"err": "unauthorized"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		}
	})

	pubGroup.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"firstname": message,
				"lastname":  nick,
			},
		})
	})

	///welcome?firstname=Jane&lastname=Doe
	myGroup.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		access_token, _ := c.Get("access_token")

		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"access_token": access_token,
				"firstname":    firstname,
				"lastname":     lastname,
			},
		})
	})

	myGroup.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")

		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"name":   name,
				"action": action,
			},
		})
	})
}
