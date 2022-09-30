package main

import (
	"fmt"
	"github.com/aiwen/aiwen-go-redis/redisConn"
)

func main() {
	fmt.Println("hello world")

	redisConn.ConnectRedis()
}
