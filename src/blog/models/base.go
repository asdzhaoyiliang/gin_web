package models

import "blog/dao"

func Count(value interface{}) int {
	var total int
	dao.DB.Model(value).Count(&total)
	return total
}
