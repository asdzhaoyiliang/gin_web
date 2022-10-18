package routers

import (
	"blog/controller"
	"blog/logger"
	"blog/settings"
	"blog/util"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	//告诉gin框架模板文件引用的静态文件
	r.Static("/static", "static")
	//告诉gin框架模板文件
	r.LoadHTMLGlob("templates/*/*")

	//中间件
	r.Use(logger.GinLogger())
	//集成session
	util.InitSession(r)

	//注册函数
	///hello函数
	r.GET("/hello", helloHandler)

	//前端

	//后端
	v2Group := r.Group("admin")
	v2Group.Use(controller.AuthMiddleware())
	admin := controller.AdminController{}
	{
		//主页面登录
		v2Group.GET("/login", admin.Login)
		v2Group.POST("/login", admin.Login)
		v2Group.GET("/logout", admin.Logout)

		//主页
		v2Group.GET("/main", admin.Main)

		//系统配置
		///页面展示
		v2Group.GET("/config", admin.Config)
		///提交更新
		v2Group.POST("/addconfig", admin.AddConfig)

		//博文列表
		v2Group.GET("/index", admin.Index)

		//博文添加
		///显示
		v2Group.GET("/article", admin.Article)
		///文章保存
		v2Group.GET("/save", admin.Save)
		///文章删除
		v2Group.GET("/delete", admin.PostDel)

		//类目
		///主页
		v2Group.GET("/category", admin.Category)
		///增加
		v2Group.GET("/categoryadd", admin.CategoryAdd)
		///保存
		v2Group.POST("/categorysave", admin.CategorySave)
		///删除
		v2Group.GET("/categorydel", admin.CategoryDel)

	}
	return r
}

func helloHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"mode": settings.Conf.Mode,
	})
}
