package models

import "time"

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
