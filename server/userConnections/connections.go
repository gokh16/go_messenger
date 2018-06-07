package userConnections

import (
	"net"

	"github.com/gorilla/websocket"
)

type Connections struct {
	//WSConnectionsMutex  *sync.Mutex
	WSConnections       map[*websocket.Conn]string // connection:login
	//TCPConnectionsMutex *sync.Mutex
	TCPConnections      map[net.Conn]string // connection:login
	OutChan             chan *Message
}

func (c *Connections) AddTCPConn(conn net.Conn, userName string, outChan *Message) *Connections {
	str := c
	str.TCPConnections[conn] = userName
	str.OutChan <- outChan
	return str
}

func (c *Connections) AddWSConn(conn *websocket.Conn, userName string, outChan *Message) *Connections {
	str := c
	str.WSConnections[conn] = userName
	str.OutChan <- outChan
	return str
}
func (c *Connections) GetAllTCPConnections() map[net.Conn]string {
	//c.TCPConnectionsMutex.Lock()
	//defer c.TCPConnectionsMutex.Unlock()
	return c.TCPConnections
}

func (c *Connections) GetAllWSConnections() map[*websocket.Conn]string {
	//c.WSConnectionsMutex.Lock()
	//defer c.WSConnectionsMutex.Unlock()
	return c.WSConnections
}
