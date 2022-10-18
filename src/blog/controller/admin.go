package controller

import (
	"blog/models"
	"blog/util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"html/template"
	"net/http"
	"strconv"
	"time"
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
		fmt.Println(password)

		//2、数据库验证
		user, _ := models.Login(username, password)
		if len(user) == 0 {
			fmt.Println("user not exist")
			c.HTML(http.StatusOK, "login.html", nil)
			return
		}

		//3、用户信息保存到session
		sess, _ := json.Marshal(user[0])
		err := util.SetSess(c, "user", sess)
		if err != nil {
			fmt.Println("login error")
			c.HTML(http.StatusOK, "login.html", nil)
			return
		}

		//4、跳转主页面
		ts := time.Now().Unix()
		url := fmt.Sprintf("/admin/main?ts=%d", ts)
		c.Redirect(http.StatusMovedPermanently, url)

	} else {
		//登录页面的操作内容
		c.HTML(http.StatusOK, "login.html", nil)
	}
}

//退出
func (a *AdminController) Logout(c *gin.Context) {
	fmt.Println("logout")
	util.DeploySess(c, "user")
	c.Redirect(http.StatusMovedPermanently, "/admin/login")
}

//注册
func (a *AdminController) Register(c *gin.Context) {

}

//主页
func (a *AdminController) Main(c *gin.Context) {
	ts, _ := c.GetQuery("ts")
	zap.L().Info("main func log...")
	c.HTML(http.StatusOK, "main.tpl", gin.H{
		"ts": ts,
	})
}

//系统配置信息展示
func (a *AdminController) Config(c *gin.Context) {
	res, _ := models.ConfigList()

	options := make(map[string]string)
	for _, v := range res {
		options[v.Name] = v.Value
	}

	datalist := gin.H{}
	datalist["config"] = options
	fmt.Println(datalist)
	c.HTML(http.StatusOK, "config.html", datalist)
}

//系统配置信息更新
func (a *AdminController) AddConfig(c *gin.Context) {
	options := make(map[string]string)
	mp := make(map[string]*models.Config)

	result, _ := models.ConfigList()
	for _, v := range result {
		options[v.Name] = v.Value
		mp[v.Name] = v
	}

	keys := []string{"url", "title", "keywords", "email", "start", "description", "qq"}
	for _, key := range keys {
		val := c.PostForm(key)
		fmt.Println("key:", key)

		if _, ok := mp[key]; !ok {
			//不存在，新建
			fmt.Println("1:", ok)
			options[key] = val
			models.UpdateConfig(&models.Config{Name: key, Value: val})
		} else {
			//存在
			fmt.Println("2:", ok)
			opt := mp[key]
			if err := models.UpdateConfig(&models.Config{Id: opt.Id, Name: key, Value: val}); err != nil {
				continue
			}
		}
	}
	msg := "success"
	info := "<script> alert('" + msg + "');window.history.go(-1); </script>"
	c.Writer.WriteString(info)

	c.Redirect(http.StatusMovedPermanently, "/admin/")

}

func AdminList(c *gin.Context) {

}

//后台首页
func (a *AdminController) Index(c *gin.Context) {
	//类目
	datalist := gin.H{}
	category, _ := models.CategoryList()
	datalist["categorys"] = category

	//商品
	var pagesize int = 5
	var offset int
	pageStr, _ := c.GetQuery("page")
	page, _ := strconv.Atoi(pageStr)
	offset = page * pagesize
	list, _ := models.GetArticleList(offset, pagesize)
	datalist["list"] = list

	//总数量
	count := models.Count(new(models.Post))
	datalist["count"] = count

	//分页
	pagebar := util.NewPager(page, count, pagesize, c.Request.RequestURI, true).ToString()
	datalist["pagebar"] = template.HTML(pagebar)

	fmt.Println(datalist)
	c.HTML(http.StatusOK, "list.html", datalist)
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
	categorylist, _ := models.CategoryList()

	datalist := gin.H{}
	datalist["categorys"] = categorylist
	c.HTML(http.StatusOK, "category.html", datalist)
}

//类目增加
func (a *AdminController) CategoryAdd(c *gin.Context) {
	idString, _ := c.GetQuery("id")
	data := gin.H{}
	fmt.Println("idString:", idString)
	if idString != "" {
		id, _ := strconv.Atoi(idString)
		category, err := models.GetCategoryById(id)
		fmt.Println("data:", err)

		if err == nil {
			data["cate"] = category[0]
		}
	}
	fmt.Println("data:", data)
	c.HTML(http.StatusOK, "category_add.html", data)
}

//类目保存
func (a *AdminController) CategorySave(c *gin.Context) {
	fmt.Println("CategorySave")

	name := c.PostForm("name")
	idString := c.PostForm("id")

	var category models.Category
	curTime := time.Now()
	if idString == "" {
		category.Name = name
		category.Created = curTime
		category.Updated = curTime
	} else {
		id, _ := strconv.Atoi(idString)
		curCategory, err := models.GetCategoryById(id)

		if err == nil {
			category.Id = id
			category.Name = name
			category.Created = curCategory[0].Created
			category.Updated = time.Now()
		}

	}
	err := models.CategoryAdd(&category)
	if err != nil {
		fmt.Println("CategorySave err :", err)
	}
	c.Redirect(http.StatusMovedPermanently, "/admin/category")

}

//类目删除
func (a *AdminController) CategoryDel(c *gin.Context) {
	idString, _ := c.GetQuery("id")
	id, _ := strconv.Atoi(idString)
	models.CategoryDel(id)
	c.Redirect(http.StatusMovedPermanently, "/admin/category")
}
