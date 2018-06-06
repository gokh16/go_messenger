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


func (c *Connections) AddTCPConn(conn net.Conn, userName string, outChan *Message) *Connections{
	str := c
	str.TCPConnections[conn] = userName
	str.OutChan <- outChan
	return str

}