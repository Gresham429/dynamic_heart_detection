package controller

import (
	"dynamic_heart_rates_detection/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type connectRequest struct {
	Url       string `json:"url"`
	Position  string `json:"position"`
	Connected bool   `json:"connected"`
}

// ConnectDevice - 设备连接
func ConnectDevice(c echo.Context) error {
	requestDevice := new(connectRequest)
	if err := c.Bind(requestDevice); err != nil {
		return c.JSON(http.StatusBadRequest, Response{Error: "Invalid JSON type"})
	}

	// 设备连接错误
	if !requestDevice.Connected {
		return c.JSON(http.StatusBadRequest, Response{Error: "设备连接错误"})
	}

	// 检查设备是否存在
	existingDevice, _ := model.GetDeviceInfo(requestDevice.Url)

	if existingDevice != nil {
		// 如果存在，则更新设备信息
		existingDevice.Connected = requestDevice.Connected
		existingDevice.Position = requestDevice.Position

		model.UpdateDevice(existingDevice)
	} else {
		// 如果不存在就新建一条记录
		device := new(model.Device)
		device.Connected = requestDevice.Connected
		device.Url = requestDevice.Url
		device.Position = requestDevice.Position

		model.CreateDevice(device)
	}

	return c.JSON(http.StatusCreated, Response{Message: "连接成功"})
}

type disconnectRequest struct {
	Url       string `json:"url"`
	Position  string `json:"position"`
	Connected bool   `json:"connected"`
}

func DisconnectDevice(c echo.Context) error {
	requestDevice := new(disconnectRequest)
	if err := c.Bind(requestDevice); err != nil {
		return c.JSON(http.StatusBadRequest, Response{Error: "Invalid JSON type"})
	}

	if requestDevice.Connected {
		return c.JSON(http.StatusBadRequest, Response{Error: "设备状态错误"})
	}

	// 检查设备是否存在
	existingDevice, _ := model.GetDeviceInfo(requestDevice.Url)

	if existingDevice != nil {
		// 如果存在，则更新设备信息
		existingDevice.Connected = requestDevice.Connected
		existingDevice.Position = requestDevice.Position

		model.UpdateDevice(existingDevice)
	} else {
		return c.JSON(http.StatusBadRequest, Response{Error: "设备不存在"})
	}

	return c.NoContent(http.StatusNoContent)
}

func GetDeviceInfo(c echo.Context) error {
	devices, err := model.GetAllDevices()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Error: "读取设备信息失败"})
	}

	return c.JSON(http.StatusOK, Response{Data: devices})
}
