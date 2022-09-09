package search

import "github.com/gin-gonic/gin"

func LoadSearch(e *gin.Engine) {
	e.GET("/search", searchHandler)
}
