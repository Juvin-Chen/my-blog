package router

import (
	"blog_backend/conf"
	"blog_backend/internal/controller"
	"blog_backend/internal/repository"
	"blog_backend/internal/service"
	"blog_backend/pkg/response"

	"github.com/gin-gonic/gin"
)

// SetupRouter 初始化 Gin 路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 1. 初始化依赖 (依赖注入)
	articleRepo := repository.NewArticleRepository(conf.DB)
	articleService := service.NewArticleService(articleRepo)
	articleCtrl := controller.NewArticleController(articleService)

	// 2. 注册测试接口
	r.GET("/ping", func(c *gin.Context) {
		response.Success(c, "pong")
	})

	// 3. 注册 API 路由组
	apiV1 := r.Group("/api/v1")
	{
		// 文章模块
		apiV1.POST("/articles", articleCtrl.Create)
		apiV1.GET("/articles", articleCtrl.List)
		apiV1.GET("/articles/:id", articleCtrl.Get)
		apiV1.PUT("/articles/:id", articleCtrl.Update)
		apiV1.DELETE("/articles/:id", articleCtrl.Delete)
	}

	return r
}
