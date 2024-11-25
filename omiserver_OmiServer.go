package omiserver

import (
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omiserd-v1"
)

type OmiServer struct {
	ServerName     string
	Address        string
	ServerRegister omiserd.Register
	HandleFuncs    map[string]func(w http.ResponseWriter, r *http.Request)
}

func newOmiServer(opts *redis.Options, serverName, address string) *OmiServer {
	return &OmiServer{
		ServerName:     serverName,
		Address:        address,
		ServerRegister: *omiserd.NewClient(opts, omiserd.Server).NewRegister(serverName, address),
		HandleFuncs:    make(map[string]func(w http.ResponseWriter, r *http.Request)),
	}
}

func (server *OmiServer) AddHanldFunc(url string, handFunc func(w http.ResponseWriter, r *http.Request)) {
	server.HandleFuncs[url] = handFunc
}

func (server *OmiServer) initHandleFunc() {
	for url, handleFunc := range server.HandleFuncs {
		http.HandleFunc(url, handleFunc)
	}
}

func (server *OmiServer) Start(weight int) {
	server.ServerRegister.RegisterAndServe(weight, func(port string) {
		server.initHandleFunc()
		err := http.ListenAndServe(port, nil)
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	})
}
