package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

func WsStatus(w http.ResponseWriter, r *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("create web socket error")
		io.WriteString(w, `{"status":"FAIL"}`)
		return
	}
	SocketStatus(ws)

}

func SocketStatus(conn *websocket.Conn) {
	for {
		messageType, _, _ := conn.ReadMessage()
		for {
			var sources []Source
			if os.Getenv("TEST") == "true" {
				sources = Select(dtbs)
			} else {
				sources = Select(NewDbConn())
			}

			jsonData, err := json.Marshal(sources)
			if err != nil {
				log.Println("error while marshall json")
				return
			}
			data := []byte(jsonData)
			if err := conn.WriteMessage(messageType, data); err != nil {
				log.Println("error while sending message")
				return
			}
			time.Sleep(30 * time.Second)
		}

	}

}

func Add(w http.ResponseWriter, r *http.Request) {

	var source Source
	err := json.NewDecoder(r.Body).Decode(&source)
	if err != nil {
		fmt.Println("Error decoding json")
		io.WriteString(w, `{"status":"FAIL"}`)
		return
	}
	if os.Getenv("TEST") == "true" {
		io.WriteString(w, Insert(dtbs, source))
	} else {
		io.WriteString(w, Insert(NewDbConn(), source))
	}

	return

}

func DeleteRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var source Source
	err := json.NewDecoder(r.Body).Decode(&source)
	if err != nil {
		fmt.Println("Error decoding json")
		io.WriteString(w, `{"status":"FAIL"}`)
		return
	}
	if os.Getenv("TEST") == "true" {
		io.WriteString(w, Delete(dtbs, source))
	} else {
		io.WriteString(w, Delete(NewDbConn(), source))
	}

	return

}
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var user User
	var username string
	var password string
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Error decoding json")
		io.WriteString(w, `{"authenticated":"false"}`)
		return
	}
	if os.Getenv("USERNAME") == "" {
		username = "admin"
	} else {
		username = os.Getenv("USERNAME")
	}
	if os.Getenv("PASSWORD") == "" {
		password = "admin"
	} else {
		password = os.Getenv("PASSWORD")
	}
	if username == user.Username && password == user.Password {
		io.WriteString(w, `{"authenticated":"true"}`)
		return
	} else {
		io.WriteString(w, `{"authenticated":"false"}`)
		return
	}

}
