package models

import (
	"blog/dao"
	"time"
)

type Comment struct {
	Id       int
	Username string
	Content  string
	Created  time.Time
	PostId   int
	Ip       string
}

func GetCommentById(post_id int) (dataList []*Comment, err error) {
	db := dao.DB.Where("post_id = ?", post_id).Order("created desc")
	err = db.Find(&dataList).Error
	return
}

func CreateComment(comment *Comment) (err error) {
	err = dao.DB.Create(&comment).Error
	return
}
