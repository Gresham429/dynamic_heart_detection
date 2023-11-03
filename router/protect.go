package router

import (
	"dynamic_heart_rates_detection/controller"

	"github.com/labstack/echo/v4"
)

func InitProtect(g *echo.Group) {
	g.GET("/get_user_info", controller.GetUserInfo)
	g.DELETE("/delete_user", controller.DeleteUser)
	g.PUT("/update_user_info", controller.UpdateUserInfo)
	g.GET("/get_device_info", controller.GetDeviceInfo)
}
