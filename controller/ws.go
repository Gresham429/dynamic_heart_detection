package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func HeartRate(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Print(err)
		return err
	}
	defer conn.Close()

	clients[conn] = true // 将新连接的客户端加入到映射表中

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return err
		}

		// 将消息广播给所有连接的客户端
		broadcast(messageType, p)
	}
}

var (
	clients = make(map[*websocket.Conn]bool)
)

func broadcast(messageType int, message []byte) {
	for client := range clients {
		err := client.WriteMessage(messageType, message)
		if err != nil {
			log.Println(err)
			client.Close()
			delete(clients, client)
		}
	}
}
