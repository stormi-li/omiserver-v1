package main

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"github.com/stormi-li/omiserver-v1"
)

var redisAddr = "118.25.196.166:3934"
var password = "12982397StrongPassw0rd"

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有来源
	},
}

func main() {
	c := omiserver.NewClient(&redis.Options{Addr: redisAddr, Password: password})
	server := c.NewOmiServer("websocket_hello", "118.25.196.166:8880")
	server.AddHanldFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// 升级 HTTP 请求到 WebSocket
		conn, err := upgrader.Upgrade(w, r, nil)
		if err == nil {
			message := fmt.Sprint("Hello ", r.URL.Query().Get("name"), ", welcome to use omi, send by websocket server")
			conn.WriteMessage(websocket.TextMessage, []byte(message))
			return
		}
		defer conn.Close()
	})
	server.Listen(1)
}
