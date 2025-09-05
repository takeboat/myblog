package model

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"column:username;not null;uniqueIndex" json:"username"`
	Password  string    `gorm:"column:password;not null" json:"password"`
	Nickname  string    `gorm:"column:nickname" json:"nickname"`
	Email     string    `gorm:"column:email" json:"email"`
	Avatar    string    `gorm:"column:avatar;default:'https://via.placeholder.com/100'" json:"avatar"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

type UserModel interface {
	Insert(user *User) error
	FindByID(id int64) (*User, error)
	FindByUsername(username string) (*User, error)
	FindByEmail(email string) (*User, error)
	Update(user *User) error
	Delete(id int64) error
}

type userModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) UserModel {
	return &userModel{db: db}
}

func (m *userModel) Insert(user *User) error {
	return m.db.Create(user).Error
}

func (m *userModel) FindByID(id int64) (*User, error) {
	var user User
	err := m.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *userModel) FindByUsername(username string) (*User, error) {
	var user User
	err := m.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (m *userModel) FindByEmail(email string) (*User, error) {
	var user User
	err := m.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *userModel) Update(user *User) error {
	return m.db.Save(user).Error
}

func (m *userModel) Delete(id int64) error {
	return m.db.Delete(&User{}, id).Error
}
