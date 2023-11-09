package middleware

import (
	"dynamic_heart_rates_detection/auth"
	"dynamic_heart_rates_detection/model"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, err := auth.GetJwtClaims(c)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, "请求头不合法")
		}

		if claims.Valid() != nil {
			return c.JSON(http.StatusUnauthorized, "请求头不合法")
		}

		if claims.ExpiresAt < time.Now().Unix() {
			return c.JSON(http.StatusUnauthorized, "请求头已过期")
		}

		user, err := model.GetUserInfo(claims.UserName)

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.JSON(http.StatusUnauthorized, "该 JWT Token 对应的用户已不存在")
			}
			return c.JSON(http.StatusInternalServerError, "数据库查询错误")
		}

		c.Set("user", &model.User{
			ID:          user.ID,
			UserName:    user.UserName,
			FullName:    user.FullName,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			Address:     user.Address,
		})

		return next(c)
	}
}
