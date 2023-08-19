package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func main() {
	router := gin.Default()

	// 定义WebSocket升级器
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// 处理WebSocket请求
	router.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 在这里处理WebSocket连接
		go handleWebSocket(conn)
	})

	router.Run(":8080")
}

func handleWebSocket(conn *websocket.Conn) {
	defer conn.Close()

	for {
		// 读取客户端发送的消息
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		// 处理接收到的消息
		// ...

		// 发送消息给客户端
		err = conn.WriteMessage(websocket.TextMessage, []byte("Hello, client!"))
		if err != nil {
			break
		}
	}
}
