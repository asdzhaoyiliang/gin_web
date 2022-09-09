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
	DB.Set("gorm:table_options", "ENGINE=innoDB DEFAULT CHARSET=utf8").AutoMigrate(&AccountInfo{})
}

func Close() {
	DB.Close()
}

//定义数据模型
type AccountInfo struct {
	gorm.Model
	Name     string `gorm:"not null;unique"`
	Password string `gorm:"not null;"`
	Status   uint   `gorm:"default:0"`
}

//ORM封装接口
type AccountInfoAPI struct {
}

func (h *AccountInfoAPI) List(offset, limit int) (accountInfo []AccountInfo) {
	return
}

func (h *AccountInfoAPI) Create(accountInfo *AccountInfo) error {
	return nil
}
func (h *AccountInfoAPI) Get(id int) (accountInfo AccountInfo) {
	return
}

func (h *AccountInfoAPI) Update(id int, updates *AccountInfo) error {
	return nil
}

func (h *AccountInfoAPI) Delete(id int) error {
	return nil
}
func (h *AccountInfoAPI) Count() (int, error) {
	var count int
	return count, nil
}
