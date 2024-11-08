package main

import (
	"log"
	"net/http"
)

/*
WebSocket是一种在单个TCP连接上进行全双工通讯的协议
。Go语言中有多个第三方库支持WebSocket，其中最流行的是gorilla/websocket。
*/
import (
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}
	defer ws.Close()

	for {
		// 读取消息
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}
		log.Printf("Received message: %s\n", p)

		// 发送消息
		err = ws.WriteMessage(messageType, p)
		if err != nil {
			log.Println("Error writing message:", err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/ws", handler)
	log.Println("Starting server at :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Println("Error starting server:", err)
	}
}
