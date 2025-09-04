package model

import "time"

// Comment 评论模型
type Comment struct {
	ID        int64      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PostID    int64      `gorm:"column:post_id;not null" json:"post_id"`
	UserID    *int64     `gorm:"column:user_id" json:"user_id"`
	Content   string     `gorm:"column:content;not null" json:"content"`
	ParentID  *int64     `gorm:"column:parent_id;comment:'父评论ID，实现嵌套'" json:"parent_id"`
	CreatedAt time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`

	// 关联关系
	Post     Post      `gorm:"foreignKey:PostID" json:"post"`
	User     *User     `gorm:"foreignKey:UserID" json:"user"`
	Parent   *Comment  `gorm:"foreignKey:ParentID" json:"parent"`
	Children []Comment `gorm:"foreignKey:ParentID" json:"children"`
}

func (Comment) TableName() string {
	return "comments"
}
