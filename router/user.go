package router

import (
	"dynamic_heart_rates_detection/controller"
	m "dynamic_heart_rates_detection/middleware"

	"github.com/labstack/echo/v4"
)

func InitUser(g *echo.Group) {
	g.POST("/register", controller.Register)
	g.POST("/login", controller.Login)
	g.POST("/register_email", controller.RegisterEmail)
	g.POST("login_with_email", controller.LoginWithEmail)
	g.GET("/get_user_info", controller.GetUserInfo, m.JwtMiddleware)
	g.DELETE("/delete_user", controller.DeleteUser, m.JwtMiddleware)
	g.PUT("/update_user_info", controller.UpdateUserInfo, m.JwtMiddleware)
}
