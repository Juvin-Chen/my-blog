package router

import (
	"blog_backend/pkg/response"

	"github.com/gin-gonic/gin"
)

// SetupRouter 初始化 Gin 路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 注册测试接口
	r.GET("/ping", func(c *gin.Context) {
		response.Success(c, "pong")
	})

	return r
}
