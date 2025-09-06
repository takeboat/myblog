package model

import (
	"blog/api/internal/types"
	"sort"
	"time"

	"gorm.io/gorm"
)

// Post 文章模型
type Post struct {
	ID         int64      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Title      string     `gorm:"column:title;not null" json:"title"`
	Content    string     `gorm:"column:content;not null" json:"content"`
	UserID     int64      `gorm:"column:user_id;not null" json:"user_id"`
	CategoryID *int64     `gorm:"column:category_id" json:"category_id"`
	Status     int        `gorm:"column:status;default:1;comment:'1: published, 0: draft'" json:"status"`
	ViewCount  int        `gorm:"column:view_count;default:0" json:"view_count"`
	CreatedAt  time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt  *time.Time `gorm:"column:deleted_at" json:"deleted_at"`

	// 关联关系
	User     User      `gorm:"foreignKey:UserID" json:"user"`
	Category *Category `gorm:"foreignKey:CategoryID" json:"category"`
	Tags     []Tag     `gorm:"many2many:post_tags;" json:"tags"`
}

func (Post) TableName() string {
	return "posts"
}

type PostTag struct {
	PostID int64 `gorm:"column:post_id;not null" json:"post_id"`
	TagID  int64 `gorm:"column:tag_id;not null" json:"tag_id"`
}

func (PostTag) TableName() string {
	return "post_tags"
}

type PostModel interface {
	Insert(post *Post) error
	FindByID(id int64) (*Post, error)
	List(page, pageSize int, categoryID *int64) ([]Post, int64, error)
	Update(post *Post) error
	Delete(id int64) error
	GetPostWithTags(id int64) (*Post, error)
	IncreaseViewCount(id int64) error
	BeginTx() *gorm.DB
	UpdateWithTx(tx *gorm.DB, post *Post) error
	RemoveTagsWithTx(tx *gorm.DB, postID int64) error
	AddTagsWithTx(tx *gorm.DB, postID int64, tagIDs []int64) error
	GetArchivedPosts(page, pageSize int) ([]types.Archive, int64, error)
}

type postModel struct {
	db *gorm.DB
}

func NewPostModel(db *gorm.DB) PostModel {
	return &postModel{db: db}
}

func (m *postModel) Insert(post *Post) error {
	return m.db.Create(post).Error
}

func (m *postModel) FindByID(id int64) (*Post, error) {
	var post Post
	err := m.db.Where("id = ? AND deleted_at IS NULL", id).First(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (m *postModel) List(page, pageSize int, categoryID *int64) ([]Post, int64, error) {
	var posts []Post
	var total int64

	query := m.db.Model(&Post{}).Where("deleted_at IS NULL")

	if categoryID != nil {
		query = query.Where("category_id = ?", *categoryID)
	}

	// 获取总数
	query.Count(&total)

	// 获取分页数据
	offset := (page - 1) * pageSize
	err := query.Offset(offset).Limit(pageSize).Find(&posts).Error
	if err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

func (m *postModel) Update(post *Post) error {
	return m.db.Save(post).Error
}

func (m *postModel) Delete(id int64) error {
	return m.db.Delete(&Post{}, id).Error
}

func (m *postModel) GetPostWithTags(id int64) (*Post, error) {
	var post Post
	err := m.db.Preload("Tags").Preload("User").Preload("Category").
		Where("id = ? AND deleted_at IS NULL", id).First(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (m *postModel) IncreaseViewCount(id int64) error {
	return m.db.Model(&Post{}).Where("id = ?", id).
		Update("view_count", gorm.Expr("view_count + 1")).Error
}

func (m *postModel) BeginTx() *gorm.DB {
	return m.db.Begin()
}
func (m *postModel) UpdateWithTx(tx *gorm.DB, post *Post) error {
	return tx.Save(post).Error
}
func (m *postModel) RemoveTagsWithTx(tx *gorm.DB, postID int64) error {
	return tx.Where("post_id = ?", postID).Delete(&PostTag{}).Error
}

func (m *postModel) AddTagsWithTx(tx *gorm.DB, postID int64, tagIDs []int64) error {
	if len(tagIDs) == 0 {
		return nil
	}
	var postTags []PostTag
	for _, tagID := range tagIDs {
		postTags = append(postTags, PostTag{
			PostID: postID,
			TagID:  tagID,
		})
	}
	return tx.Create(&postTags).Error
}

func (m *postModel) GetArchivedPosts(page, pageSize int) ([]types.Archive, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	// 获取文章总数
	var total int64
	err := m.db.Model(&Post{}).Where("deleted_at IS NULL and status = 1").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	// 分页获取文章归档列表
	offset := (page - 1) * pageSize
	var posts []Post
	err = m.db.Model(&Post{}).Where("deleted_at IS NULL and status = 1").Order("created_at DESC").
		Offset(offset).Limit(pageSize).Find(&posts).Error
	if err != nil {
		return nil, 0, err
	}
	// 将文章按照 年 月 分开
	var mapArchive = make(map[string][]types.PostArchive)
	for _, post := range posts {
		yearMonth := post.CreatedAt.Format("2006-01")
		postArchive := types.PostArchive{
			Id:        post.ID,
			Title:     post.Title,
			CreatedAt: post.CreatedAt.Format("2006-01-02"),
		}
		mapArchive[yearMonth] = append(mapArchive[yearMonth], postArchive)
	}
	// 将其聚合
	var archives []types.Archive
	for key, postarchives := range mapArchive {
		var t time.Time
		t, _ = time.Parse("2006-01", key)
		archives = append(archives, types.Archive{
			Year:  t.Year(),
			Month: int(t.Month()),
			Posts: postarchives,
		})
	}
	// 按照时间排序
	sort.Slice(archives, func(i, j int) bool {
		if archives[i].Year != archives[j].Year {
			return archives[i].Year > archives[j].Year
		}
		return archives[i].Month > archives[j].Month
	})
	return archives, total, nil
}
