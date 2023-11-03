package main

import (
	m "dynamic_heart_rates_detection/middleware"
	"dynamic_heart_rates_detection/model"
	"dynamic_heart_rates_detection/router"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func ConnectDatabase() {
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

func StartWeb() {
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

	p := gAPI.Group("/protected")

	// 配置 JWT 的中间件
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(m.JwtCustomClaims)
		},
		SigningKey: []byte("Gresham"),
	}

	p.Use(echojwt.WithConfig(config))

	router.InitProtect(p)

	log.Fatal(e.Start(":1323"))
}

func main() {
	// 连接数据库
	ConnectDatabase()

	// 启动 Web
	StartWeb()

	// 关闭数据库连接
	defer model.CloseDatabase()
}
