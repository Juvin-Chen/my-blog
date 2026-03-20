package main

import (
	"blog_backend/conf"
	"blog_backend/internal/router"
	"fmt"
)

func main() {
	// 1. 初始化数据库配置
	// 格式: 用户名:密码@tcp(地址:端口)/数据库名?charset=utf8mb4&parseTime=True&loc=Local
	dsn := "root:root@tcp(127.0.0.1:3306)/blog_db?charset=utf8mb4&parseTime=True&loc=Local"
	conf.InitDB(dsn)

	// 2. 初始化路由
	r := router.SetupRouter()

	// 3. 启动服务
	port := ":8080"
	fmt.Printf("Blog backend server is running at http://localhost%s\n", port)
	if err := r.Run(port); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
