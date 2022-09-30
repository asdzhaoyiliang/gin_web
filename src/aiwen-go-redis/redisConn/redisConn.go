package redisConn

import (
	"github.com/aiwen/aiwen-go-redis/config"
	"github.com/go-redis/redis/v8"
	"log"
)

func ConnectRedis() *redis.Client {
	conn := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
	if _, err := conn.Ping(conn.Context()).Result(); err != nil {
		log.Fatal("redis connect err:", err.Error())
	}
	return conn
}
