package routers

import "github.com/gin-gonic/gin"

func shopHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "shop",
	})
}

func LoadShop(e *gin.Engine) {
	e.GET("/shop", shopHandler)
}
