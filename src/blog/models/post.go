package models

import (
	"blog/dao"
	"time"
)

type Post struct {
	Id         int
	UserId     int
	Title      string
	Url        string
	Content    string
	Tags       string
	Views      int
	IsTop      int
	Created    time.Time
	Updated    time.Time
	CategoryId int
	Status     int
	Types      int
	Info       string
	Image      string
}

func GetArticleList(offset int, pagesize int) (articleList []*Post, err error) {
	db := dao.DB.Offset(offset).Limit(pagesize).Order(" is_top desc, created desc")
	if err = db.Find(&articleList).Error; err != nil {
		return nil, err
	}

	return articleList, nil
}

func GetDetailById(id int) (post *Post, err error) {
	post = new(Post)
	db := dao.DB.Where("id = ?", id).First(&post)
	if err = db.Error; err != nil {
		return nil, err
	}

	return post, nil
}

func CreatePost(post *Post) (err error) {
	return dao.DB.Create(&post).Error
}

func UpdatePost(post *Post) (err error) {
	return dao.DB.Save(post).Error
}
func DeletePost(id int) (err error) {
	return dao.DB.Where("id = ?", id).Delete(&Post{}).Error
}

func GetAllArticle(keyword string, cate_id int, actionName string, page int, pageSize int, sort_type int) (articleList []*Post, err error) {
	db := dao.DB
	if cate_id > 0 {
		db = db.Where("category_id = ?", cate_id)
	}
	if len(keyword) > 0 {
		db = db.Where("title like ?", "%"+keyword+"%")
	}
	if actionName == "resource" {
		db = db.Where("types = ?", 0)
	} else {
		db = db.Where("types = ?", 1)
	}
	if actionName == "home" {
		db = db.Where("is_top = ?", 1)
	}
	offset := (page - 1) * pageSize
	db = db.Offset(offset).Limit(pageSize).Order("views desc")

	if sort_type == 1 {
		db = db.Order("views desc")
	} else {
		db = db.Order("created desc")
	}
	if err = db.Find(&articleList).Error; err != nil {
		return nil, err
	}
	return articleList, nil
}

func GetNotice(cate_id int) (articleList []*Post, err error) {
	db := dao.DB.Where("category_id = ?", cate_id)
	if err = db.Find(&articleList).Error; err != nil {
		return nil, err
	}
	return
}
