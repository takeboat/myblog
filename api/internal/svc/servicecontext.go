package svc

import (
	"blog/api/internal/config"
	"blog/api/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config        config.Config

	CategoryModel model.CategoryModel
	PostModel     model.PostModel
	TagModel      model.TagModel
	UserModel     model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DataSource), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:        c,
		CategoryModel: model.NewCategoryModel(db),
		PostModel:     model.NewPostModel(db),
		TagModel:      model.NewTagModel(db),
		UserModel:     model.NewUserModel(db),
	}
}
