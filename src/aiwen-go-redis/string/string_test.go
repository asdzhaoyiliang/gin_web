package string

import (
	"fmt"
	"github.com/aiwen/aiwen-go-redis/redisConn"
	"testing"
)

var conn = redisConn.ConnectRedis()
var ctx = conn.Context()

func TestSet(t *testing.T) {
	conn.Set(ctx, "name", "z2", 0)
}

func TestGet(t *testing.T) {
	val := conn.Get(ctx, "name").Val()
	fmt.Println(val)
}
func TestIncr(t *testing.T) {
	conn.Incr(ctx, "inc")
}
