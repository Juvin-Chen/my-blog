package model

import "gorm.io/gorm"

// User 管理员模型
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null;type:varchar(100)" json:"username"`
	Password string `gorm:"not null;type:varchar(100)" json:"-"` // JSON 返回时忽略密码
}

// Category 分类模型
type Category struct {
	gorm.Model
	Name     string    `gorm:"unique;not null;type:varchar(100)" json:"name"`
	Articles []Article `gorm:"foreignKey:CategoryID" json:"articles,omitempty"`
}

// Tag 标签模型
type Tag struct {
	gorm.Model
	Name     string    `gorm:"unique;not null;type:varchar(100)" json:"name"`
	Articles []Article `gorm:"many2many:article_tags;" json:"articles,omitempty"`
}

// Article 文章模型
type Article struct {
	gorm.Model
	Title      string   `gorm:"type:varchar(255);not null" json:"title"`
	Content    string   `gorm:"type:longtext;not null" json:"content"`
	Desc       string   `gorm:"type:varchar(500)" json:"desc"`
	CategoryID uint     `json:"category_id"`
	Category   Category `gorm:"foreignKey:CategoryID" json:"category"`
	Tags       []Tag    `gorm:"many2many:article_tags;" json:"tags"`
}
