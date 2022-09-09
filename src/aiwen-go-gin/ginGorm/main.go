package main

import (
	"fmt"
	"github.com/aiwen/aiwen-go-gin/ginGorm/sql"
	"github.com/gin-gonic/gin"
)

//采用的技术架构
//gin：web服务框架
//gorm:ORM
//mysql:数据库
func main() {
	//打开数据库
	sql2.Init()
	defer sql2.Close()
	//api操作接口
	//gin框架web服务
	////gin路由创建
	r := gin.Default()
	////函数注册
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "hello")
	})
	r.POST("/add", addHandler)
	r.GET("/list", listHandler)
	r.GET("/update/:id", updateHandler)
	r.GET("/delete/:id", deleteHandler)
	r.GET("/count", countHandler)
	////服务启动
	if err := r.Run(); err != nil {
		fmt.Println("startup failed,err:%v\n", err)
	}
}

func addHandler(c *gin.Context) {

}
func listHandler(c *gin.Context) {

}
func updateHandler(c *gin.Context) {

}
func deleteHandler(c *gin.Context) {

}
func countHandler(c *gin.Context) {

}
