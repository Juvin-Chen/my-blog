package repository

import (
	"blog_backend/internal/model"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	BaseRepository
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

// Create 创建文章
func (r *ArticleRepository) Create(article *model.Article) error {
	return r.db.Create(article).Error
}

// GetByID 根据ID获取文章
func (r *ArticleRepository) GetByID(id uint) (*model.Article, error) {
	var article model.Article
	err := r.db.Preload("Category").Preload("Tags").First(&article, id).Error
	return &article, err
}

// List 分页获取文章列表
func (r *ArticleRepository) List(page, size int) ([]model.Article, int64, error) {
	var articles []model.Article
	var total int64

	db := r.db.Model(&model.Article{})
	db.Count(&total)

	offset := (page - 1) * size
	err := db.Preload("Category").Preload("Tags").
		Offset(offset).Limit(size).
		Order("created_at DESC").
		Find(&articles).Error

	return articles, total, err
}

// Update 更新文章
func (r *ArticleRepository) Update(article *model.Article) error {
	// 使用 Save 会更新所有字段，包括关联关系（many2many）
	return r.db.Save(article).Error
}

// Delete 删除文章
func (r *ArticleRepository) Delete(id uint) error {
	return r.db.Delete(&model.Article{}, id).Error
}
