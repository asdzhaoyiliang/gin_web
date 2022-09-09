package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func LoggerWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "123456")
		log.Print("before")
		c.Next()
		log.Print("after")
		latency := time.Since(t)
		log.Print(latency)
		status := c.Writer.Status()
		log.Print(status)
	}
}
func main() {
	r := gin.Default()

	r.Use(LoggerWare())
	r.GET("/middleware", middleware)
	r.GET("/middleware2", middleware)

	if err := r.Run(); err != nil {
		fmt.Println("err:", err)
	}
}

func middleware(c *gin.Context) {
	example := c.GetString("example")
	log.Print(example)
}
