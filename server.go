package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func showLog(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
}

func main() {
	http.HandleFunc("/", showLog)
	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, req *http.Request) { w.Write([]byte("ok")) })
	http.ListenAndServe(":3000", nil)
}
