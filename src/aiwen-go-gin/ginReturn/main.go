package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/returnjson", returnJson)
	r.GET("/returnresp/:name/:age", returnResp)

	if err := r.Run(); err != nil {
		fmt.Println("err:", err)
	}
}

type Person struct {
	Name string `form:"name"`
	Age  int    `form:"age"`
}

func returnJson(c *gin.Context) {
	var person Person
	err := c.ShouldBindQuery(&person)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"name": person.Name,
		"age":  person.Age,
	})
}

type Response struct {
	Code    int
	Message string
	Data    interface{}
}

type Student struct {
	Name string `uri:"name"`
	Age  int    `uri:"age"`
}

func returnResp(c *gin.Context) {
	var json Student
	err := c.ShouldBindUri(&json)
	if err != nil {
		c.JSON(400, map[string]interface{}{"message": err.Error()})
		return
	}
	var res Response
	res.Code = 200
	res.Data = json
	res.Message = "ok"
	c.JSON(200, res)
}
