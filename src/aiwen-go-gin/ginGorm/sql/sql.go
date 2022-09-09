package sql2

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

//定义全局变量
var DB *gorm.DB
var err error

//创建mysql连接
func Init() {

	DB, err = gorm.Open("mysql", "wanart:wanart@/ginsql?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(nil)
	}
	DB.Set("gorm:table_options", "ENGINE=innoDB DEFAULT CHARSET=utf8").AutoMigrate(&User{})
}

func Close() {
	DB.Close()
}

//定义数据模型
type User struct {
	gorm.Model
	Name     string `gorm:"not null;unique"`
	Password string `gorm:"not null;"`
	Status   uint   `gorm:"default:0"`
}

//ORM封装接口
