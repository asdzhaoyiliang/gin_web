package routers

import (
	"blog/controller"
	"blog/logger"
	"blog/settings"
	"blog/util"
	"github.com/gin-gonic/gin"
	"strings"
)

func GetReqName() gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.Request.RequestURI
		var actionName = "home"
		if strings.HasSuffix(uri, "resource") {
			c.Set("actionName", "resource")
			actionName = "resource"
		} else if strings.HasSuffix(uri, "article") {
			actionName = "article"
		} else if uri == "/admin/login" {
			actionName = "login"
			c.Set("test", "admin")
		}
		c.Set("actionName", actionName)
		c.Next()
	}
}
func SetRouter() *gin.Engine {
	r := gin.Default()

	//告诉gin框架模板文件引用的静态文件
	r.Static("/static", "static")
	//告诉gin框架模板文件
	r.LoadHTMLGlob("templates/*/*")

	//中间件
	r.Use(logger.GinLogger())
	r.Use(GetReqName())
	//集成session
	util.InitSession(r)

	//注册函数
	///hello函数
	r.GET("/hello", helloHandler)

	//前端
	v1Group := r.Group("blog")
	blog := controller.BlogController{}
	{
		v1Group.GET("/home", blog.GetHome)
		v1Group.GET("/article", blog.GetArticleList)
		v1Group.GET("/detail/:id", blog.GetArticleDetail)
		v1Group.POST("/comment", blog.CreateComment)
		v1Group.GET("/resource", blog.Resource)
		v1Group.GET("/about", blog.GetAbout)
	}

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
		///显示详情
		v2Group.GET("/article", admin.Article)
		///文章保存
		v2Group.POST("/save", admin.Save)
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
