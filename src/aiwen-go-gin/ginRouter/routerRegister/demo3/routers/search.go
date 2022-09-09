package routers

import "github.com/gin-gonic/gin"

func searchHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "search",
	})
}
func LoadSearch(e *gin.Engine) {
	e.GET("/search", searchHandler)
}
