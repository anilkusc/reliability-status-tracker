package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	_ "github.com/mattn/go-sqlite3"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var restart = false
var dtbs *sql.DB

func main() {
	dtbs = NewDbConn()
	go Control()
	r := mux.NewRouter()

	r.HandleFunc("/status", WsStatus)
	r.HandleFunc("/add/", Add).Methods("POST")
	r.HandleFunc("/delete/", DeleteRecord).Methods("POST")
	r.HandleFunc("/login/", Login).Methods("POST")
	fmt.Println("Serving on:8080")
	http.ListenAndServe(":8080", r)
}
