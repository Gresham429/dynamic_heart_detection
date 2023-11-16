package model

import (
	"context"
	"dynamic_heart_rates_detection/config"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB = nil
var RDB *redis.Client = nil

// 初始化 postgres
func initPostgres() error {
	dbConfig := config.JsonConfiguration.DB
	dsn := "host=" + dbConfig.Host + " user=" + dbConfig.User + " password=" + dbConfig.Password + " dbname=" + dbConfig.DBName + " port=" + dbConfig.Port + " sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	// 自动迁移模型
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Device{})

	DB = db

	return nil
}

// 初始化 redis
func initRedis() error {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "gresham_rdb_1:6379",
		Password: "20040420", // no password set
		DB:       0,          // use default DB
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}

	RDB = rdb

	return nil
}

// 连接 postgres 与 redis
func ConnectDatabase() {
	var err error

	for DB == nil {
		err = initPostgres()
		if err != nil {
			log.Printf("Error connecting to the postgres: %v\n", err)
			log.Println("Retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
		}
	}

	log.Println("Connecte postgres successfully!")

	for RDB == nil {
		err = initRedis()
		if err != nil {
			log.Printf("Error connecting to the redis: %v\n", err)
			log.Println("Retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
		}
	}

	log.Println("Connecte redis successfully!")
}

// 关闭 postgres 与 redis
func CloseDatabase() {
	if DB != nil {
		db, err := DB.DB()
		if err != nil {
			fmt.Println("Error getting underlying database:", err)
		}

		// 关闭 postgres 连接
		if err := db.Close(); err != nil {
			fmt.Println("Error closing postgres:", err)
		}
	}

	if RDB != nil {
		// 关闭 redis 连接
		if err := RDB.Close(); err != nil {
			fmt.Println("Error closing redis:", err)
		}
	}
}
