package shop

import "github.com/gin-gonic/gin"

func shopHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "shop4",
	})
}
