package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

//定义注册函数
//w:响应内容
//r:请求内容
func sayHelloName(w http.ResponseWriter, r *http.Request) {
	//业务逻辑
	r.ParseForm()       //解析参数，默认是不解析的
	fmt.Println(r.Form) //请求参数<k,v>数据对
	fmt.Println("path = ", r.URL.Path)

	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello World") //这个写入w输出到客户端
}
func main() {

	//通过net/http包注册函数
	http.HandleFunc("/hello", sayHelloName) // 设置访问的路由/hello
	//设置监听端口
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	//http://localhost:8080/hello?name=zyl&age=18
}
