package model

import (
	"dynamic_heart_rates_detection/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() (*gorm.DB, error) {
	dbConfig := config.JsonConfiguration.DB
	dsn := "host=" + dbConfig.Host + " user=" + dbConfig.User + " password=" + dbConfig.Password + " dbname=" + dbConfig.DBName + " port=" + dbConfig.Port + " sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	// 自动迁移模型
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Device{})

	DB = db

	return DB, nil
}

func CloseDatabase() {
	if DB != nil {
		db, err := DB.DB()
		if err != nil {
			fmt.Println("Error getting underlying database:", err)
		}

		// 关闭数据库连接
		if err := db.Close(); err != nil {
			fmt.Println("Error closing database:", err)
		}
	}
}
