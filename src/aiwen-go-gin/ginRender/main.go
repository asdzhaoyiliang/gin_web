package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/byte", byteFmtHandler)
	r.GET("/str", strFmtHandler)

	if err := r.Run(); err != nil {
		fmt.Println("err:", err)
	}
}
func byteFmtHandler(c *gin.Context) {
	fullPath := c.FullPath()
	fmt.Println("fullPath:", fullPath)
	c.Writer.Write([]byte(fullPath))
}
func strFmtHandler(c *gin.Context) {
	fullpath := c.FullPath()
	fmt.Println("fullPath:", fullpath)
	c.Writer.WriteString(fullpath)
}
