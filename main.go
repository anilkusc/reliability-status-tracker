package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var restart = false

func main() {
	//go Control()
	r := mux.NewRouter()

	r.HandleFunc("/status", WsStatus)
	r.HandleFunc("/add/", Add).Methods("POST")
	r.HandleFunc("/delete/", DeleteRecord).Methods("POST")
	r.HandleFunc("/login/", Login).Methods("POST")
	r.HandleFunc("/", Hello).Methods("POST")
	fmt.Println("Serving on:8080")
	http.ListenAndServe(":8080", r)
	/*http.HandleFunc("/status", WsStatus)
	http.HandleFunc("/add", Add)
	http.ListenAndServe(":8080", nil)
	*/
}
