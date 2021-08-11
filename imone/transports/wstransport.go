package transports

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"imdemo/imone/services"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},

}

func serveWs(service services.WSService,w http.ResponseWriter,r *http.Request)  {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	//注册客户端
	//client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	//client.hub.register <- client
	//
	//
	//go client.writePump()
	//go client.readPump()
	client := &services.Client{
		UserId:   strconv.Itoa(rand.Intn(100)),
		Address: ":8080",
		Conn:    conn,
		Send:    make(chan []byte, 256),
	}

	err=service.RegisterClient(client)
	if err != nil {
		log.Println(err)
		return
	}
	go client.WritePump()
	go client.ReadPump()

}


func MakeHttpHandler(service services.WSService)http.Handler  {
	r:= mux.NewRouter()
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(service, w, r)
		
	})
	return r

}