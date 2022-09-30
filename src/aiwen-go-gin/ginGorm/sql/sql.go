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
	DB.Offset(offset).Limit(limit).Find(&accountInfo)
	return accountInfo
}

func (h *AccountInfoAPI) Create(accountInfo *AccountInfo) error {
	err := DB.Create(accountInfo).Error
	return err
}
func (h *AccountInfoAPI) Get(p string) (accountInfo AccountInfo) {
	DB.Where("password=?", p).First(&accountInfo)
	return accountInfo
}

func (h *AccountInfoAPI) Update(id int, updates *AccountInfo) error {
	var accountInfo AccountInfo
	err := DB.First(&accountInfo, id).Error
	if err != nil {
		return err
	}
	err = DB.Model(&accountInfo).Update(updates).Error
	return err
}

func (h *AccountInfoAPI) Delete(id int) error {
	var accountInfo AccountInfo
	err := DB.First(&accountInfo, id).Error
	if err != nil {
		return err
	}
	err = DB.Delete(&accountInfo).Error
	return err
}
func (h *AccountInfoAPI) Count() (int, error) {
	var count int
	return count, nil
}
