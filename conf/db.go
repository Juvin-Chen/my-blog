package conf

import (
	"blog_backend/internal/model"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB 初始化 MySQL 连接并自动迁移模型
func InitDB(dsn string) {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Database connection successful")

	// 自动迁移模型
	err = DB.AutoMigrate(
		&model.User{},
		&model.Category{},
		&model.Tag{},
		&model.Article{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	fmt.Println("Database migration successful")
}
