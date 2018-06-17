package userConnections

import (
	"net"
	"sync"
	"github.com/gorilla/websocket"
	"go_messenger/server/service/serviceModels"
)

type Connections struct {
	WSConnectionsMutex  *sync.Mutex
	WSConnections       map[*websocket.Conn]string // connection:login
	TCPConnectionsMutex *sync.Mutex
	TCPConnections      map[net.Conn]string // connection:login
	OutChan             chan *serviceModels.MessageOut
}

func InitConnections() *Connections {
	instance := Connections{}
	instance.WSConnectionsMutex = new(sync.Mutex)
	instance.WSConnections = make(map[*websocket.Conn]string, 0)
	instance.TCPConnectionsMutex = new(sync.Mutex)
	instance.TCPConnections = make(map[net.Conn]string, 0)
	instance.OutChan = make(chan *serviceModels.MessageOut, 1024)
	return &instance
}

func (c *Connections) AddTCPConn(conn net.Conn, userName string) {
	c.TCPConnectionsMutex.Lock()
	c.TCPConnections[conn] = userName
	c.TCPConnectionsMutex.Unlock()
}

func (c *Connections) AddWSConn(conn *websocket.Conn, userName string) {
	c.WSConnectionsMutex.Lock()
	c.WSConnections[conn] = userName
	c.WSConnectionsMutex.Unlock()
}

func (c *Connections) GetAllTCPConnections() map[net.Conn]string {
	return c.TCPConnections
}

func (c *Connections) GetAllWSConnections() map[*websocket.Conn]string {
	return c.WSConnections
}
