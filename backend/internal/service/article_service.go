package service

import (
	"blog_backend/internal/model"
	"blog_backend/internal/repository"
)

// ArticleService 文章服务
type ArticleService struct {
	BaseService
	repo *repository.ArticleRepository
}

func NewArticleService(repo *repository.ArticleRepository) *ArticleService {
	return &ArticleService{repo: repo}
}

// CreateArticle 创建文章
func (s *ArticleService) CreateArticle(article *model.Article) error {
	return s.repo.Create(article)
}

// GetArticle 获取文章详情
func (s *ArticleService) GetArticle(id uint) (*model.Article, error) {
	return s.repo.GetByID(id)
}

// GetArticleList 获取文章列表
func (s *ArticleService) GetArticleList(page, size int) ([]model.Article, int64, error) {
	return s.repo.List(page, size)
}

// UpdateArticle 更新文章
func (s *ArticleService) UpdateArticle(id uint, data map[string]interface{}) error {
	// 1. 先查询文章是否存在
	article, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	// 2. 这里的逻辑可以根据实际业务复杂化
	// 如果是全量更新：
	if title, ok := data["title"].(string); ok {
		article.Title = title
	}
	if content, ok := data["content"].(string); ok {
		article.Content = content
	}
	if desc, ok := data["desc"].(string); ok {
		article.Desc = desc
	}
	if categoryID, ok := data["category_id"].(float64); ok { // JSON 解析出的数字默认为 float64
		article.CategoryID = uint(categoryID)
	}

	// 处理标签更新（简化逻辑，实际可能需要更复杂的关联处理）
	if tagIDs, ok := data["tag_ids"].([]interface{}); ok {
		var tags []model.Tag
		for _, tid := range tagIDs {
			if idFloat, ok := tid.(float64); ok {
				tag := model.Tag{}
				tag.ID = uint(idFloat)
				tags = append(tags, tag)
			}
		}
		article.Tags = tags
	}

	return s.repo.Update(article)
}

// DeleteArticle 删除文章
func (s *ArticleService) DeleteArticle(id uint) error {
	return s.repo.Delete(id)
}
