package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//实例化对象，对应yaml文件
var Conf = new(AppConfig)

//AppConfig struct
type AppConfig struct {
	Mode         string `mapstructure:"mode"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
}

//Mysql config struct
type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"db"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

//LogConfig
type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

//定义一个初始化函数
//1、读取config.yaml
//2、文件变更，自动识别（main.go web 热加载自动）
func Init() error {
	//目的：读取-》conf

	viper.SetConfigFile("./conf/config.yaml")
	//viper.SetConfigFile("E:/gin_web/src/blog/conf/config.yaml")

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err := viper.Unmarshal(&Conf); err != nil {
			panic(fmt.Errorf("Unmarshal failed,err:%v", err))
		}
	})
	viper.WatchConfig()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("ReadConfig failed,err:%v", err))
	}
	if err = viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("Unmarshal failed,err:%v", err))
	}
	return nil
}
