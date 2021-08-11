package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

	"log"
	"net/http"
)




var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},

}


func serveWs(hub *Hub,w http.ResponseWriter,r *http.Request)  {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	//注册客户端
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client


	go client.writePump()
	go client.readPump()

}



func main() {
	hub := NewHub()
	go hub.run()
	r:=mux.NewRouter()
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	http.ListenAndServe(":8080",r)
}