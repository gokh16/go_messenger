package userConnections

import (
	"github.com/gorilla/websocket"
	"net"
)

type Connections struct{
	WSConnections map[*websocket.Conn]string // connection:login
	TCPConnections map[net.Conn]string // connection:login
	OutChan chan *Message
}

func (c *Connections) AddTCPConn() *Connections{

}