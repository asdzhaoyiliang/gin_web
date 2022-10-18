package models

import (
	"blog/dao"
	"time"
)

type Category struct {
	Id      int
	Name    string
	Created time.Time
	Updated time.Time
}

func CategoryAdd(category *Category) (err error) {
	err = dao.DB.Save(category).Error
	return err
}

func CategoryDel(id int) error {
	return dao.DB.Where("id = ?", id).Delete(&Category{}).Error
}

func CategoryUpdate(category *Category) error {
	return dao.DB.Save(category).Error
}

func CategoryList() (categoryList []*Category, err error) {
	err = dao.DB.Find(&categoryList).Error
	if err != nil {
		return nil, err
	}
	return categoryList, nil
}

func GetCategoryById(id int) (category []*Category, err error) {
	err = dao.DB.Find(&category, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}
