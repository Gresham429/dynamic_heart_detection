package router

import (
	"dynamic_heart_rates_detection/controller"

	"github.com/labstack/echo/v4"
)

func InitUser(g *echo.Group) {
	g.POST("/register", controller.Register)
	g.POST("/login", controller.Login)
}
