package model

import (
	"time"

	"gorm.io/gorm"
)

// Category 分类模型
type Category struct {
	ID        int64      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name      string     `gorm:"column:name;not null;uniqueIndex" json:"name"`
	CreatedAt time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (Category) TableName() string {
	return "categories"
}

type CategoryModel interface {
	Insert(category *Category) error
	FindByID(id int64) (*Category, error)
	FindByName(name string) (*Category, error)
	List() ([]Category, error)
	Update(category *Category) error
	Delete(id int64) error
}

type categoryModel struct {
	db *gorm.DB
}

func NewCategoryModel(db *gorm.DB) CategoryModel {
	return &categoryModel{db: db}
}

func (m *categoryModel) Insert(category *Category) error {
	return m.db.Create(category).Error
}

func (m *categoryModel) FindByID(id int64) (*Category, error) {
	var category Category
	err := m.db.Where("id = ?", id).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (m *categoryModel) FindByName(name string) (*Category, error) {
	var category Category
	err := m.db.Where("name = ?", name).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (m *categoryModel) List() ([]Category, error) {
	var categories []Category
	err := m.db.Where("deleted_at IS NULL").Find(&categories).Error
	return categories, err
}

func (m *categoryModel) Update(category *Category) error {
	return m.db.Save(category).Error
}

func (m *categoryModel) Delete(id int64) error {
	return m.db.Delete(&Category{}, id).Error
}
