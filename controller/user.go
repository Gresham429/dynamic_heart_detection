package controller

import (
	"dynamic_heart_rates_detection/auth"
	"dynamic_heart_rates_detection/model"
	"dynamic_heart_rates_detection/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// Register - 用户注册
func Register(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid JSON type",
		})
	}

	// 检查用户是否已经存在
	existingUser, _ := model.GetUserInfo(user.UserName)
	if existingUser != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "用户名已存在",
		})
	}

	// 对密码进行哈希处理
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashPassword)

	// 创建用户
	if err := model.CreateUser(user); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "注册成功",
	})
}

// Login - 用户登录(生成JWT令牌)
func Login(c echo.Context) error {
	loginUser := new(model.User)
	if err := c.Bind(loginUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request",
		})
	}

	// 检索用户信息
	user, _ := model.GetUserInfo(loginUser.UserName)
	if user == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "用户名或密码错误",
		})
	}

	// 核对密码信息
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "用户名或密码错误",
		})
	}

	// 生成 jwt 令牌
	jwt, err := auth.GenerateJWTToken(user.UserName)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"token": jwt})
}

// GetUser - 获取用户信息（需要JWT身份验证）
func GetUserInfo(c echo.Context) error {
	user := c.Get("user")

	userInfo := user.(*model.User)

	response := map[string]interface{}{
		"id":           strconv.FormatUint(uint64(userInfo.ID), 10),
		"username":     userInfo.UserName,
		"full_name":    utils.GetValueOrEmptyString(userInfo.FullName),
		"email":        utils.GetValueOrEmptyString(userInfo.Email),
		"phone_number": utils.GetValueOrEmptyString(userInfo.PhoneNumber),
		"address":      utils.GetValueOrEmptyString(userInfo.Address),
	}

	return c.JSON(http.StatusOK, response)
}

// UpdateUserInfo - 更新用户信息（需要JWT身份验证）
func UpdateUserInfo(c echo.Context) error {
	user := c.Get("user")

	userInfo := user.(*model.User)

	// 从请求中获得需要更新的用户信息
	updatedInfo := new(model.User)
	if err := c.Bind(updatedInfo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid JSON type",
		})
	}

	// 当 JSON 中存在以下信息之一时，更新 user
	if updatedInfo.Password != "" {
		// 对密码进行哈希处理
		HashPassword, err := bcrypt.GenerateFromPassword([]byte(updatedInfo.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		userInfo.Password = string(HashPassword)
	}
	if updatedInfo.FullName != nil {
		userInfo.FullName = updatedInfo.FullName
	}
	if updatedInfo.Email != nil {
		userInfo.Email = updatedInfo.Email
	}
	if updatedInfo.PhoneNumber != nil {
		userInfo.PhoneNumber = updatedInfo.PhoneNumber
	}
	if updatedInfo.Address != nil {
		userInfo.Address = updatedInfo.Address
	}

	// Save the updated user information to the database
	if err := model.UpdateUser(userInfo); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "无法更新用户信息",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "成功更新用户信息"})
}

// DeleteUser - 删除用户（需要JWT身份验证）
func DeleteUser(c echo.Context) error {
	user := c.Get("user")

	userInfo := user.(*model.User)

	err := model.DeleteUser(userInfo.UserName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "删除用户失败"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "删除用户成功"})
}
