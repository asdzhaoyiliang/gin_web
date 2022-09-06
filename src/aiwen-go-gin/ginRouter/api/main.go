package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	r := gin.Default()

	//http://127.0.0.1:8080/hello/zyl/play
	r.GET("/user/:name/*action", func(c *gin.Context) {
		//业务逻辑
		name := c.Param("name")
		action := c.Param("action")
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+"\t"+action)
	})
	r.Run(":8080")
}
