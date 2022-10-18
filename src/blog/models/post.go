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
