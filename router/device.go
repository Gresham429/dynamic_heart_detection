package router

import (
	"dynamic_heart_rates_detection/controller"

	"github.com/labstack/echo/v4"
)

func InitDevice(g *echo.Group) {
	g.POST("/connect", controller.ConnectDevice)
	g.POST("/disconnect", controller.DisconnectDevice)
}
