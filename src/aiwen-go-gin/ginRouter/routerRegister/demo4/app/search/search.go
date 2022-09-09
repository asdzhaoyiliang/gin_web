package search

import "github.com/gin-gonic/gin"

func searchHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "search4",
	})
}
