package settings

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	if err := Init(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(Conf.MysqlConfig.Host)
	fmt.Println(Conf.Mode)
}
