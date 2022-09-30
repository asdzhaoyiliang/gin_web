package main

import (
	"fmt"
	"github.com/aiwen/aiwen-go-redis/redisConn"
	"github.com/go-redis/redis/v8"
)

var api = redisConn.ConnectRedis()

func main() {
	SetZSort()
	GetZSort()
}

func SetZSort() {
	ctx := api.Context()
	api.ZAdd(ctx, "weibo", &redis.Z{
		Member: "a",
		Score:  1,
	})
	api.ZAdd(ctx, "weibo", &redis.Z{
		Member: "b",
		Score:  6,
	})
	api.ZAdd(ctx, "weibo", &redis.Z{
		Member: "c",
		Score:  3,
	})
	api.ZAdd(ctx, "weibo", &redis.Z{
		Member: "d",
		Score:  2,
	})
}

func GetZSort() {
	ctx := api.Context()
	data := api.ZRevRangeWithScores(ctx, "weibo", 0, -1)
	for k, v := range data.Val() {
		fmt.Println(v.Member, "    ", v.Score, "       ", k)
	}
}
