package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var account = gin.H{
	"z1": gin.H{
		"name": "z1",
		"age":  21,
	},
	"z2": gin.H{
		"name": "z2",
		"age":  22,
	},
}

func main() {
	r := gin.Default()

	basicAut := gin.BasicAuth(gin.Accounts{
		"z1": "1",
		"z2": "2",
		"z3": "3",
	})
	//按照先后顺序执行basicAut，basicAuth
	r.GET("auth", basicAut, basicAuth)

	if err := r.Run(); err != nil {
		fmt.Println(err)
	}
}

func basicAuth(c *gin.Context) {
	name := c.GetString(gin.AuthUserKey)
	if user, ok := account[name]; ok {
		c.JSON(200, gin.H{
			"user": user,
			"name": name,
		})
	} else {
		c.JSON(400, gin.H{
			"user": user,
			"name": name,
		})
	}
}
