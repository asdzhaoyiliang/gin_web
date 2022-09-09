package routers

import (
	"github.com/gin-gonic/gin"
)

type Option func(e *gin.Engine)

var option = []Option{}

func Include(opts ...Option) {
	option = append(option, opts...)
}
func Init() *gin.Engine {
	r := gin.Default()
	for _, v := range option {
		v(r)
	}
	return r
}
