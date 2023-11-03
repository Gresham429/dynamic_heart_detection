package controller

import (
	"dynamic_heart_rates_detection/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ConnectDevice - 设备连接
func ConnectDevice(c echo.Context) error {
	device := new(model.Device)
	if err := c.Bind(device); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid JSON type",
		})
	}

	// 设备连接错误
	if !device.Connected {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "设备连接错误",
		})
	}

	// 检查设备是否存在
	ExistingDevice, _ := model.GetDeviceInfo(device.Url)

	if ExistingDevice != nil {
		// 如果存在，则更新设备信息
		model.UpdateDevice(device)
	} else {
		// 如果不存在就新建一条记录
		model.CreateDevice(device)
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "连接成功",
	})
}

func DisconnectDevice(c echo.Context) error {
	device := new(model.Device)
	if err := c.Bind(device); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid JSON type",
		})
	}

	if device.Connected {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "设备状态错误",
		})
	}

	// 检查设备是否存在
	ExistingDevice, _ := model.GetDeviceInfo(device.Url)

	if ExistingDevice != nil {
		// 如果存在，则更新设备信息
		model.UpdateDevice(device)
	} else {
		return c.JSON(http.StatusCreated, map[string]string{
			"message": "设备不存在",
		})
	}

	return c.NoContent(http.StatusNoContent)
}

func GetDeviceInfo(c echo.Context) error {
	devices, err := model.GetAllDevices()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "读取设备信息失败",
		})
	}

	return c.JSON(http.StatusOK, devices)
}