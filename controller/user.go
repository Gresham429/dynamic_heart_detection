package controller

import (
	"dynamic_heart_rates_detection/auth"
	"dynamic_heart_rates_detection/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type registerRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// Register - 用户注册
func Register(c echo.Context) error {
	registerUser := new(registerRequest)
	if err := c.Bind(registerUser); err != nil {
		return c.JSON(http.StatusBadRequest, Response{Error: "Invalid JSON type"})
	}

	// 检查用户是否已经存在
	existingUser, err := model.GetUserInfo(registerUser.UserName)
	if existingUser != nil {
		return c.JSON(http.StatusBadRequest, Response{Error: "用户名已存在"})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Error: err.Error()})
	}

	if registerUser.UserName == "" || registerUser.Password == "" {
		return c.JSON(http.StatusInternalServerError, Response{Error: "用户名和密码不可以为空"})
	}

	// 对密码进行哈希处理
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(registerUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	registerUser.Password = string(hashPassword)

	user := &model.User{
		UserName: registerUser.UserName,
		Password: registerUser.Password,
	}

	// 创建用户
	if err := model.CreateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Error: err.Error()})
	}

	return c.JSON(http.StatusCreated, Response{Message: "注册成功"})
}

type loginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}

// Login - 用户登录(生成JWT令牌)
func Login(c echo.Context) error {
	loginUser := new(loginRequest)
	if err := c.Bind(loginUser); err != nil {
		return c.JSON(http.StatusBadRequest, Response{Error: "Invalid request"})
	}

	// 检索用户信息
	user, err := model.GetUserInfo(loginUser.UserName)
	if user == nil {
		return c.JSON(http.StatusUnauthorized, Response{Error: "用户名或密码错误"})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Error: err.Error()})
	}

	// 核对密码信息
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, Response{Error: "用户名或密码错误"})
	}

	// 生成 jwt 令牌
	jwt, err := auth.GenerateJWTToken(user.UserName)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, Response{Error: err.Error()})
	}

	response := loginResponse{Token: jwt}

	return c.JSON(http.StatusOK, Response{Data: response})
}

type userInfoResponse struct {
	UserName string `json:"username,omitempty"`
	FullName string `json:"full_name,omitempty"`
	Email    string `json:"email,omitempty"`
	Address  string `json:"address,omitempty"`
}

// GetUser - 获取用户信息（需要JWT身份验证）
func GetUserInfo(c echo.Context) error {
	userName, ok := c.Get("username").(string)

	if !ok {
		// 类型断言失败，处理错误
		return c.JSON(http.StatusInternalServerError, Response{Error: "无法将 user_name 转换为字符串"})
	}

	userInfo, err := model.GetUserInfo(userName)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, Response{Error: err.Error()})
	}

	response := userInfoResponse{
		UserName: userInfo.UserName,
		FullName: userInfo.FullName,
		Email:    userInfo.Email,
		Address:  userInfo.Address,
	}

	return c.JSON(http.StatusOK, Response{Data: response})
}

type updateRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

// UpdateUserInfo - 更新用户信息（需要JWT身份验证）
func UpdateUserInfo(c echo.Context) error {
	userName, ok := c.Get("username").(string)

	if !ok {
		// 类型断言失败，处理错误
		return c.JSON(http.StatusInternalServerError, Response{Error: "无法将 user_name 转换为字符串"})
	}

	userInfo, err := model.GetUserInfo(userName)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	// 从请求中获得需要更新的用户信息
	updatedInfo := new(updateRequest)
	if err := c.Bind(updatedInfo); err != nil {
		return c.JSON(http.StatusBadRequest, Response{Error: "Invalid JSON type"})
	}

	// 当 JSON 中存在以下信息之一时，更新 user
	if updatedInfo.Password != "" {
		// 对密码进行哈希处理
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(updatedInfo.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		userInfo.Password = string(hashPassword)
	}

	if updatedInfo.UserName != "" {
		userInfo.UserName = updatedInfo.UserName
	}

	if updatedInfo.FullName != "" {
		userInfo.FullName = updatedInfo.FullName
	}

	if updatedInfo.Email != "" {
		userInfo.Email = updatedInfo.Email
	}

	if updatedInfo.Address != "" {
		userInfo.Address = updatedInfo.Address
	}

	// Save the updated user information to the database
	if err := model.UpdateUser(userInfo); err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Error: "无法更新用户信息"})
	}

	return c.JSON(http.StatusOK, Response{Message: "成功更新用户信息"})
}

// DeleteUser - 删除用户（需要JWT身份验证）
func DeleteUser(c echo.Context) error {
	userName, ok := c.Get("username").(string)

	if !ok {
		// 类型断言失败，处理错误
		return c.JSON(http.StatusInternalServerError, Response{Error: "无法将 user_name 转换为字符串"})
	}

	err := model.DeleteUser(userName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Error: "删除用户失败"})
	}

	return c.JSON(http.StatusOK, Response{Message: "删除用户成功"})
}
