package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	fmt.Println("hello world")
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	// 3.设置监听端口(默认8080)
	r.Run(":8080")
}
