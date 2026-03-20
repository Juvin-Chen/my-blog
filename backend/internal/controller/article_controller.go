package controller

import (
	"blog_backend/internal/model"
	"blog_backend/internal/service"
	"blog_backend/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	BaseController
	service *service.ArticleService
}

func NewArticleController(service *service.ArticleService) *ArticleController {
	return &ArticleController{service: service}
}

// CreatePostArticleRequest 创建文章请求结构体
type CreatePostArticleRequest struct {
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
	Desc       string `json:"desc"`
	CategoryID uint   `json:"category_id" binding:"required"`
	TagIDs     []uint `json:"tag_ids"`
}

// Create 创建文章
// POST /api/v1/articles
func (ctrl *ArticleController) Create(c *gin.Context) {
	var req CreatePostArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid parameters: "+err.Error())
		return
	}

	// 组装模型
	article := &model.Article{
		Title:      req.Title,
		Content:    req.Content,
		Desc:       req.Desc,
		CategoryID: req.CategoryID,
	}

	// 处理标签
	for _, tid := range req.TagIDs {
		tag := model.Tag{}
		tag.ID = tid
		article.Tags = append(article.Tags, tag)
	}

	if err := ctrl.service.CreateArticle(article); err != nil {
		response.Error(c, 500, "Failed to create article")
		return
	}

	response.Success(c, article)
}

// List 获取文章列表
// GET /api/v1/articles?page=1&size=10
func (ctrl *ArticleController) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	if page <= 0 {
		page = 1
	}
	if size <= 0 || size > 100 {
		size = 10
	}

	articles, total, err := ctrl.service.GetArticleList(page, size)
	if err != nil {
		response.Error(c, 500, "Failed to fetch articles")
		return
	}

	response.Success(c, gin.H{
		"list":  articles,
		"total": total,
		"page":  page,
		"size":  size,
	})
}

// Get 获取单篇文章
// GET /api/v1/articles/:id
func (ctrl *ArticleController) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, 400, "Invalid ID format")
		return
	}

	article, err := ctrl.service.GetArticle(uint(id))
	if err != nil {
		response.Error(c, 404, "Article not found")
		return
	}

	response.Success(c, article)
}

// Update 更新文章
// PUT /api/v1/articles/:id
func (ctrl *ArticleController) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, 400, "Invalid ID format")
		return
	}

	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		response.Error(c, 400, "Invalid parameters")
		return
	}

	if err := ctrl.service.UpdateArticle(uint(id), data); err != nil {
		response.Error(c, 500, "Failed to update article")
		return
	}

	response.Success(c, nil)
}

// Delete 删除文章
// DELETE /api/v1/articles/:id
func (ctrl *ArticleController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(c, 400, "Invalid ID format")
		return
	}

	if err := ctrl.service.DeleteArticle(uint(id)); err != nil {
		response.Error(c, 500, "Failed to delete article")
		return
	}

	response.Success(c, nil)
}
