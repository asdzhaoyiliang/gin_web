package main

import (
	"github.com/aiwen/aiwen-go-gin/ginGorm/sql"
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
	////函数注册
	////服务启动
}
