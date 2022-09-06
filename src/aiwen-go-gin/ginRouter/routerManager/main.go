package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	v1 := r.Group("v1")
	v1.GET("/login", login)
	v1.GET("/register", register)

	r.Run(":8080")
}
func login(c *gin.Context) {
	//http://10.10.10.2:8080/v1/login?name=z
	name := c.DefaultQuery("name", "defaultName")
	c.String(200, "login:%s\n", name)
}
func register(c *gin.Context) {
	//http://10.10.10.2:8080/v1/register?name=z
	name := c.DefaultQuery("name", "defaultName")
	c.String(200, "register:%s\n", name)
}
