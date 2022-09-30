package controller

import (
	"blog/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminController struct {
}

//登录
func (a *AdminController) Login(c *gin.Context) {
	if c.Request.Method == "POST" {
		//username
		//password
		//submit login-> index

		//1、接受参数
		username := c.PostForm("username")
		password := c.PostForm("password")

		//2、数据库验证
		models.Login(username, password)
		//3、用户信息保存到session

		//4、跳转主页面

	} else {
		//登录页面的操作内容
		c.HTML(http.StatusOK, "login.html", nil)
	}
}

//退出
func (a *AdminController) Logout(c *gin.Context) {

}

//注册
func (a *AdminController) Register(c *gin.Context) {

}

//主页
func (a *AdminController) Main(c *gin.Context) {

}

//系统配置信息展示
func (a *AdminController) Config(c *gin.Context) {

}

//系统配置信息更新
func (a *AdminController) AddConfig(c *gin.Context) {

}

func AdminList(c *gin.Context) {

}

//后台首页
func (a *AdminController) Index(c *gin.Context) {

}

//博文添加
func (a *AdminController) Article(c *gin.Context) {

}

//保存
func (a *AdminController) Save(c *gin.Context) {

}

//文章删除
func (a *AdminController) PostDel(c *gin.Context) {

}

//类目
func (a *AdminController) Category(c *gin.Context) {

}

//类目增加
func (a *AdminController) CategoryAdd(c *gin.Context) {

}

//类目保存
func (a *AdminController) CategorySave(c *gin.Context) {

}

//类目删除
func (a *AdminController) CategoryDel(c *gin.Context) {

}
