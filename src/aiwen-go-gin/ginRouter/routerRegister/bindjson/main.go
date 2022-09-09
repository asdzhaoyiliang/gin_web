package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	Name string `form:"name"`
	Age  int    `form:"age"`
	Sex  string `form:"sex"`
}
type Person2 struct {
	Name string `form:"name"`
	Age  int    `form:"age"`
	Sex  string `form:"sex"`
}
type Student struct {
	Name string `uri:"name"`
	Id   string `uri:"id"`
}

func main() {
	r := gin.Default()

	r.POST("/bindjson", jsonHandler)
	r.GET("/shouldbindquery", shouldBindQueryHandler)
	r.GET("/shouldbinduri/:name/:id", shouldBindUriHandler)

	if err := r.Run(); err != nil {
		fmt.Println("err:", err)
	}
}
func jsonHandler(c *gin.Context) {
	//声明解析接受json的变量
	var json Person
	//解析变量
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}

	//业务处理

	//返回结果
	c.JSON(200, gin.H{
		"status": "ok",
		"name":   json.Name,
		"age":    json.Age,
		"sex":    json.Sex,
	})
}
func shouldBindUriHandler(c *gin.Context) {
	var stu Student
	name := c.Param("name")
	fmt.Println(name)
	if err := c.ShouldBindUri(&stu); err != nil {
		c.JSON(400, gin.H{"message:": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"name": stu.Name,
		"id":   stu.Id,
	})
}

func shouldBindQueryHandler(c *gin.Context) {
	var jsonw Person2

	//sex := c.Param("sex")
	//age := c.Param("age")
	if err := c.ShouldBindQuery(&jsonw); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}

	//fmt.Println("sex:%s", sex)
	//fmt.Println("age:%s", age)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		//"name":   jsonw.Name,
		"age": jsonw.Age,
		"sex": jsonw.Sex,
	})
}
