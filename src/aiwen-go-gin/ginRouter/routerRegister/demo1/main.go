package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//路由
	r := gin.Default()

	//路由注册函数
	r.GET("hello", helloHandler)

	//启动
	if err := r.Run(); err != nil {
		fmt.Println("startup failed,err:%v\n", err)
	}
}

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}
