package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/user", func(c *gin.Context) {
		name := c.DefaultQuery("name", "defaultName")
		c.String(http.StatusOK, "hello %s\n", name)

	})
	r.Run(":8080")

}
