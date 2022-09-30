package util

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

//初始化session操作
func InitSession(engine *gin.Engine) {
	store, err := redis.NewStore(10, "tcp", "127.0.0.1:6379", "", []byte("secret"))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Init Redis Session Success")
	}
	engine.Use(sessions.Sessions("sessionid", store))
}

func SetSess(context *gin.Context, key interface{}, value interface{}) error {
	session := sessions.Default(context)
	if session == nil {
		return nil
	}
	session.Set(key, value)
	return session.Save()
}
func GetSess(context *gin.Context, key interface{}) interface{} {
	session := sessions.Default(context)
	return session.Get(key)
}

func DeploySess(context *gin.Context, key interface{}) {
	session := sessions.Default(context)
	session.Set(key, nil)
	session.Save()
}
