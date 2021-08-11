package common

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net"
)

var MsgChan chan string

type Client struct {

	UserId string `json:"userId"`
	Address string  `json:"address"` //所在服务器
	Conn *websocket.Conn `json:"conn"`
	Send chan []byte	`json:"-"`
}


func GetIp() bool {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
		return false
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						fmt.Println(ipnet.IP.String())
						//return true
					}
				}
			}
		}
	}

	return false
}


