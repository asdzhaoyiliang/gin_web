package routers

import "github.com/gin-gonic/gin"

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}
func SetupRouter() *gin.Engine {
	//路由
	r := gin.Default()

	//路由注册函数
	r.GET("hello", helloHandler)
	return r
}
