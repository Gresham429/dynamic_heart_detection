package main

import (
	"dynamic_heart_rates_detection/config"
	"dynamic_heart_rates_detection/model"
	"dynamic_heart_rates_detection/router"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func connectDatabase() {
	var db *gorm.DB
	var err error

	for db == nil {
		db, err = model.InitDatabase()
		if err != nil {
			log.Printf("Error connecting to the database: %v\n", err)
			log.Println("Retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
		}
	}

	log.Println("Connecte database successfully!")
}

func startWeb() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 允许所有跨域请求
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // 或者特定的域
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	gAPI := e.Group("/api")

	gUser := gAPI.Group("/user")
	router.InitUser(gUser)

	gDevice := gAPI.Group("/device")
	router.InitDevice(gDevice)

	log.Fatal(e.Start(":" + config.JsonConfiguration.WebPort))
}

func main() {
	args := os.Args

	config.InitConfig(args[1])

	// 连接数据库
	connectDatabase()

	// 启动 Web
	startWeb()

	// 关闭数据库连接
	defer model.CloseDatabase()
}
