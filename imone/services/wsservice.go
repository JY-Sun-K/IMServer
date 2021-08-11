package services

import (
	"bytes"
	"github.com/gorilla/websocket"
	"imdemo/imone/common"
	"imdemo/imone/dao"
	"log"
	"time"
)


const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)


var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type ClientI interface {
	writePump()
	readPump()
}

type Client struct {

	UserId string
	Address string //所在服务器
	Conn *websocket.Conn
	Send chan []byte
}


func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			common.MsgChan<- string(message)
			log.Println("send message: ",message)

			// Add queued chat messages to the current websocket message.
			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) ReadPump() {
	defer func() {
		rdb:= &dao.RedisHubImpl{}
		ws :=&WSServiceImpl{rdb}
		ws.UnRegisterClient(c)
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.Send <- message
		common.MsgChan<- string(message)
		log.Println("receive message",message)
	}
}

type WSService interface {
	RegisterClient(client *Client)error
	UnRegisterClient(client *Client)error
}

type WSServiceImpl struct {
	RedisHub dao.RedisHub
}

func MakeWSServiceImpl(redisHub dao.RedisHub) WSService {
	return &WSServiceImpl{RedisHub: redisHub}
}

func (w *WSServiceImpl) RegisterClient(client *Client) error{
	//clients :=&common.Client{
	//	UserId:  client.UserId,
	//	Address: ":8080",
	//	Conn:    client.Conn,
	//	Send:    make(chan []byte, 256),
	//}
	log.Println("注册:",client)
	err := w.RedisHub.RegisterClient((*common.Client)(client))

	if err != nil {
		return err
	}
	return nil
}

func (w *WSServiceImpl) UnRegisterClient(client *Client)error {
	err := w.RedisHub.UnRegisterClient(client.UserId)
	if err != nil {
		return err
	}
	return nil
}