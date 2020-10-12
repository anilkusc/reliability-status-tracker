package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var restart = false

func main() {
	go Control()
	http.HandleFunc("/status", WsStatus)
	http.ListenAndServe(":8080", nil)

}
