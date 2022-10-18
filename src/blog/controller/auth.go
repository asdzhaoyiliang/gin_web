package controller

import (
	"blog/models"
	"blog/util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//认证登录中间件
func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		uri := c.Request.RequestURI
		fmt.Println("uri:", uri)

		if uri == "/admin/login" || uri == "/admin/logout" {
			return
		}

		//获取session
		sess := util.GetSess(c, "user")
		if sess == nil {
			//c.JSON(http.StatusBadRequest, gin.H{
			//	"errmsg": "login error",
			//})

			c.Redirect(http.StatusMovedPermanently, "/admin/login")
			c.Abort()
			return
		}

		//解析session，println
		var member models.User
		json.Unmarshal(sess.([]byte), &member)
		fmt.Println("models.user:", member.Username)
		c.Next()
	}
}
