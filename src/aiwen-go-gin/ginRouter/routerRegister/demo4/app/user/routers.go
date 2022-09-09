package user

import "github.com/gin-gonic/gin"

func LoadUser(e *gin.Engine) {
	v := e.Group("user")
	v.GET("/login", LoginHandler)
	v.GET("/register", RegisterHandler)
}
