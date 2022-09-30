package main

import (
	"fmt"
	"github.com/aiwen/aiwen-go-gin/ginGorm/sql"
	"github.com/gin-gonic/gin"
	"strconv"
)

var api = &sql2.AccountInfoAPI{}

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
	r.GET("/get/:p", getHandler)
	////服务启动
	if err := r.Run(); err != nil {
		fmt.Println("startup failed,err:%v\n", err)
	}
}

func addHandler(c *gin.Context) {
	var accountInfo sql2.AccountInfo
	if err := c.ShouldBindJSON((&accountInfo)); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	if err := api.Create(&accountInfo); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"msg":  "success",
		"data": accountInfo,
	})

}
func listHandler(c *gin.Context) {
	offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	res := api.List(offset, limit)
	c.JSON(200, gin.H{
		"msg":  "success",
		"data": res,
	})
}
func updateHandler(c *gin.Context) {
	var accountInfo sql2.AccountInfo
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.ShouldBindJSON(&accountInfo); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	res := api.Update(id, &accountInfo)
	c.JSON(200, gin.H{
		"msg":  "success",
		"data": res,
	})
}
func deleteHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	res := api.Delete(id)
	c.JSON(200, gin.H{
		"msg":  "success",
		"data": res,
	})
}
func getHandler(c *gin.Context) {
	p := c.Param("p")
	fmt.Println(p)
	res := api.Get(p)
	c.JSON(200, gin.H{
		"msg":  "success",
		"data": res,
	})
}
func countHandler(c *gin.Context) {

}
