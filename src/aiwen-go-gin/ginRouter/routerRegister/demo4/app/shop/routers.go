package shop

import "github.com/gin-gonic/gin"

func LoadShop(e *gin.Engine) {
	e.GET("/shop", shopHandler)
}
