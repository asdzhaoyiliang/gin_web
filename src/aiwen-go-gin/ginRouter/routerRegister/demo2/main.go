package main

import (
	"fmt"
	"github.com/aiwen/aiwen-go-gin/ginRouter/routerRegister/demo2/routers"
)

func main() {
	r := routers.SetupRouter()

	//启动
	if err := r.Run(); err != nil {
		fmt.Println("startup failed,err:%v\n", err)
	}
}
