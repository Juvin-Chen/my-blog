package model

import (
	"time"

	"gorm.io/gorm"
)

// Model 基础模型，直接定义字段以方便引用，同时保持与 gorm.Model 兼容
type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// User 管理员模型
type User struct {
	Model
	Username string `gorm:"unique;not null;type:varchar(100)" json:"username"`
	Password string `gorm:"not null;type:varchar(100)" json:"-"` // JSON 返回时忽略密码
}

// Category 分类模型
type Category struct {
	Model
	Name     string    `gorm:"unique;not null;type:varchar(100)" json:"name"`
	Articles []Article `gorm:"foreignKey:CategoryID" json:"articles,omitempty"`
}

// Tag 标签模型
type Tag struct {
	Model
	Name     string    `gorm:"unique;not null;type:varchar(100)" json:"name"`
	Articles []Article `gorm:"many2many:article_tags;" json:"articles,omitempty"`
}

// Article 文章模型
type Article struct {
	Model
	Title      string   `gorm:"type:varchar(255);not null" json:"title"`
	Content    string   `gorm:"type:longtext;not null" json:"content"`
	Desc       string   `gorm:"type:varchar(500)" json:"desc"`
	CategoryID uint     `json:"category_id"`
	Category   Category `gorm:"foreignKey:CategoryID" json:"category"`
	Tags       []Tag    `gorm:"many2many:article_tags;" json:"tags"`
}
