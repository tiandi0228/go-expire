package dao

import (
	"hongcha/go-expire/internal/app/model"
	"hongcha/go-expire/internal/base/db"
	"hongcha/go-expire/internal/base/mistake"
)

// GetCategoryAll 获取所有分类
func GetCategoryAll() (category []*model.Category, err error) {
	err = db.Engine.Find(&category)
	if err != nil {
		err = mistake.NewDaoErr(err)
	}
	return
}

// AddCategory 添加分类
func AddCategory(category *model.Category) (err error) {
	_, err = db.Engine.Insert(category)
	if err != nil {
		err = mistake.NewDaoErr(err)
	}
	return
}
