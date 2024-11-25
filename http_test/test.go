package main

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omiserver-v1"
)

var redisAddr = "118.25.196.166:3934"
var password = "12982397StrongPassw0rd"

func main() {
	c := omiserver.NewClient(&redis.Options{Addr: redisAddr, Password: password})
	server := c.NewOmiServer("http_hello", "118.25.196.166:8889")
	server.AddHanldFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello", r.URL.Query().Get("name"), ", welcome to use omi, send by http server")
	})
	server.Listen(1)
}
