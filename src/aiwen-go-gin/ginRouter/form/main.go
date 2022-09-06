package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/form", func(c *gin.Context) {
		//获取类型
		types := c.DefaultPostForm("username", "defaultValue")
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.String(http.StatusOK, "user: %s\npassword: %s\ntype:%s\n", username, password, types)
	})
	r.Run(":8080")
}
