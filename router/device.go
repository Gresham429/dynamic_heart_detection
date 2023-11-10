package router

import (
	"dynamic_heart_rates_detection/controller"
	m "dynamic_heart_rates_detection/middleware"

	"github.com/labstack/echo/v4"
)

func InitDevice(g *echo.Group) {
	g.POST("/connect", controller.ConnectDevice)
	g.POST("/disconnect", controller.DisconnectDevice)
	g.GET("/get_device_info", controller.GetDeviceInfo, m.JwtMiddleware)
}
