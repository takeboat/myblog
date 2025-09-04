package model

import (
	"time"

	"gorm.io/gorm"
)

// Tag 标签模型
type Tag struct {
	ID        int64      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name      string     `gorm:"column:name;not null;uniqueIndex" json:"name"`
	CreatedAt time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (Tag) TableName() string {
	return "tags"
}

type TagModel interface {
	Insert(tag *Tag) error
	FindByID(id int64) (*Tag, error)
	FindByName(name string) (*Tag, error)
	List() ([]Tag, error)
	Update(tag *Tag) error
	Delete(id int64) error
}

type tagModel struct {
	db *gorm.DB
}

func NewTagModel(db *gorm.DB) TagModel {
	return &tagModel{db: db}
}

func (m *tagModel) Insert(tag *Tag) error {
	return m.db.Create(tag).Error
}

func (m *tagModel) FindByID(id int64) (*Tag, error) {
	var tag Tag
	err := m.db.Where("id = ?", id).First(&tag).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (m *tagModel) FindByName(name string) (*Tag, error) {
	var tag Tag
	err := m.db.Where("name = ?", name).First(&tag).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (m *tagModel) List() ([]Tag, error) {
	var tags []Tag
	err := m.db.Where("deleted_at IS NULL").Find(&tags).Error
	return tags, err
}

func (m *tagModel) Update(tag *Tag) error {
	return m.db.Save(tag).Error
}

func (m *tagModel) Delete(id int64) error {
	return m.db.Delete(&Tag{}, id).Error
}
