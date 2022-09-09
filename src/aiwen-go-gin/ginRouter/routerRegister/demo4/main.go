package main

import (
	"fmt"
	"github.com/aiwen/aiwen-go-gin/ginRouter/routerRegister/demo4/app/search"
	"github.com/aiwen/aiwen-go-gin/ginRouter/routerRegister/demo4/app/shop"
	"github.com/aiwen/aiwen-go-gin/ginRouter/routerRegister/demo4/app/user"
	"github.com/aiwen/aiwen-go-gin/ginRouter/routerRegister/demo4/routers"
)

func main() {
	routers.Include(shop.LoadShop, search.LoadSearch, user.LoadUser)

	r := routers.Init()

	//启动
	if err := r.Run(); err != nil {
		fmt.Println("startup failed,err:%v\n", err)
	}
}
