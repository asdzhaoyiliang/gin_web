package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	//通过httprouter创建路由
	r := httprouter.New()
	//通过路由设置规则（注册函数）
	//http://127.0.0.1:8080/
	r.GET("/", Index)
	//http://127.0.0.1:8080/hello/zyl
	r.GET("/hello/:name", Hello)
	//启动
	err := http.ListenAndServe("10.10.10.2:8080", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, " Welcome !\n")
}
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//获取参数
	name := ps.ByName("name")
	//写入w
	fmt.Fprintf(w, "Hello, %s\n", name)
}
