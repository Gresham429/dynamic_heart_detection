package middleware

import (
	"dynamic_heart_rates_detection/auth"
	"dynamic_heart_rates_detection/config"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// 获得 JWT token string
		authorization := strings.Split(c.Request().Header.Get("Authorization"), " ")

		if len(authorization) < 2 {
			return c.JSON(http.StatusUnauthorized, "请求头不合法")
		}

		if authorization[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, "请求头不合法")
		}

		tokenString := authorization[1]

		// 验证 JWT token
		parsedToken, err := jwt.ParseWithClaims(tokenString, &auth.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JsonConfiguration.JwtSecret), nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, err)
		}

		// 获取声明
		claims, ok := parsedToken.Claims.(*auth.JwtCustomClaims)

		if !ok || !parsedToken.Valid {
			return c.JSON(http.StatusUnauthorized, "Token 不合法")
		}

		c.Set("username", claims.UserName)

		return next(c)
	}
}
