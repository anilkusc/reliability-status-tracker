package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func WsStatus(w http.ResponseWriter, r *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	SocketStatus(ws)

}

func SocketStatus(conn *websocket.Conn) {
	for {
		sources := Select()
		messageType, _, _ := conn.ReadMessage()

		for _, source := range sources {
			go func() {
				for {
					jsonData, err := json.Marshal(source)
					if err != nil {
						log.Println(err)
					}
					data := []byte(jsonData)
					if err := conn.WriteMessage(messageType, data); err != nil {
						log.Println(err)
						return
					}
					time.Sleep(30 * time.Second)
				}
			}()

		}

	}

}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Connected")
	Socket(ws)

}

func Socket(conn *websocket.Conn) {
	for {
		var source Source
		var target Target
		messageType, request, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		err = json.Unmarshal(request, &source)
		if err != nil {
			fmt.Println("error:", err)
		}
		for {
			resp, err := http.Get(source.Host)
			if err != nil {
				fmt.Println(err)
				data := []byte("error")
				conn.WriteMessage(messageType, data)
				break
			}
			defer resp.Body.Close()
			target.Host = source.Host
			target.Status = resp.StatusCode
			jsonData, err := json.Marshal(target)
			if err != nil {
				log.Println(err)
			}
			data := []byte(jsonData)
			if err := conn.WriteMessage(messageType, data); err != nil {
				log.Println(err)
				return
			}
			time.Sleep(30 * time.Second)
		}

	}
}
