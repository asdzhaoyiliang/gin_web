package routers

import "github.com/gin-gonic/gin"

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello2",
	})
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/hello", helloHandler)
	return r
}
