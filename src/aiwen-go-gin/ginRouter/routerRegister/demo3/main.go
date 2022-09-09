package main

import (
	"fmt"
	"github.com/aiwen/aiwen-go-gin/ginRouter/routerRegister/demo3/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routers.LoadShop(r)
	routers.LoadSearch(r)
	//启动
	if err := r.Run(); err != nil {
		fmt.Println("startup failed,err:%v\n", err)
	}
}
