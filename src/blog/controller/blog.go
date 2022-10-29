package controller

import (
	"blog/models"
	"blog/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type BlogController struct {
}

func list(c *gin.Context) map[string]interface{} {
	datalist := gin.H{}
	var (
		page     int
		cateId   int
		pageSize int = 3
		keyword  string
	)
	if pageStr, ok := c.GetQuery("page"); ok {
		page, _ = strconv.Atoi(pageStr)
	}

	if cateIdStr, ok := c.GetQuery("cate_id"); ok {
		cateId, _ = strconv.Atoi(cateIdStr)
	}
	keyword, _ = c.GetQuery("keyword")
	if page < 1 {
		page = 1
	}
	categoryList, _ := models.CategoryList()
	count := models.Count(new(models.Post))
	actionName := c.GetString("actionName")

	articleList, _ := models.GetAllArticle(keyword, cateId, actionName, page, pageSize, 0)
	hotList, _ := models.GetAllArticle(keyword, cateId, "", 0, 5, 0)

	pagebar := util.NewPager(page, int(count), pageSize, "/v1/article", true).ToString()
	datalist["cates"] = categoryList
	datalist["list"] = articleList
	datalist["hot"] = hotList
	datalist["pagebar"] = pagebar
	return datalist
}
func (b *BlogController) GetHome(c *gin.Context) {
	dataList := list(c)
	cate_id := 2
	notices, _ := models.GetNotice(cate_id)
	dataList["notices"] = notices
	c.HTML(http.StatusOK, "home.html", dataList)
}

func (b *BlogController) GetArticleList(c *gin.Context) {
	dataList := list(c)
	c.HTML(http.StatusOK, "article.html", dataList)
}

func (b *BlogController) GetArticleDetail(c *gin.Context) {
	var (
		id int
	)
	idString, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效id",
		})
		return
	}
	id, _ = strconv.Atoi(idString)
	article, _ := models.GetDetailById(id)
	var content = template.HTML(article.Content)

	comments, _ := models.GetCommentById(id)
	categoryList, _ := models.CategoryList()
	hotList, _ := models.GetAllArticle("", 0, "", 0, 5, 1)
	dataList := gin.H{
		"post":     article,
		"content":  content,
		"cates":    categoryList,
		"hot":      hotList,
		"comments": comments,
	}
	fmt.Println(dataList)
	c.HTML(http.StatusOK, "detail.html", dataList)
}

func (b *BlogController) CreateComment(c *gin.Context) {
	comment := models.Comment{}
	username, _ := c.GetPostForm("username")
	comment.Username = username
	content, _ := c.GetPostForm("content")
	comment.Content = content
	post_id, _ := c.GetPostForm("post_id")
	comment.PostId, _ = strconv.Atoi(post_id)
	comment.Created = time.Now()
	var msg = "发布成功"
	if err := models.CreateComment(&comment); err != nil {
		msg = "发布失败"
	}
	info := "<script>alert('" + msg + "');window.history.go(-1);</script>"
	c.Writer.WriteString(info)
	fmt.Println(post_id)
	c.Redirect(http.StatusMovedPermanently, "/blog/detail/"+post_id)
}

func (b *BlogController) Resource(c *gin.Context) {
	dataList := list(c)
	dataList["actionName"] = "resource"
	c.HTML(http.StatusOK, "resource.html", dataList)
}

func (b *BlogController) GetAbout(c *gin.Context) {
	id := 1
	article, _ := models.GetDetailById(id)
	c.HTML(http.StatusOK, "about.html", gin.H{
		"content": template.HTML(article.Content),
	})
}
