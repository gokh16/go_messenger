package userConnections

import (
	"net"
	"sync"

	"go_messenger/server/service/serviceModels"

	"github.com/gorilla/websocket"
	"log"
)

//Connections is a structure with connections and channel for write out data
type Connections struct {
	wsConnectionsMutex  *sync.Mutex
	wsConnections       map[*websocket.Conn]string // connection:login
	tcpConnectionsMutex *sync.Mutex
	tcpConnections      map[net.Conn]string // connection:login
	OutChan             chan *serviceModels.MessageOut
}

//Function InitConnections is init for Connections struct
func InitConnections() *Connections {
	instance := Connections{}
	instance.wsConnectionsMutex = new(sync.Mutex)
	instance.wsConnections = make(map[*websocket.Conn]string, 0)
	instance.tcpConnectionsMutex = new(sync.Mutex)
	instance.tcpConnections = make(map[net.Conn]string, 0)
	instance.OutChan = make(chan *serviceModels.MessageOut, 1024)
	return &instance
}

//AddTCPConn method is adding incoming connection with login to source structure
func (c *Connections) AddTCPConn(conn net.Conn, userName string) {
	c.tcpConnectionsMutex.Lock()
	c.tcpConnections[conn] = userName
	c.tcpConnectionsMutex.Unlock()
	log.Println(c.tcpConnections[conn], "ADDTCP")
}

//DeleteTCPConn method is deleting suitable connection from tcpConnections map of Connections struct
func (c *Connections) DeleteTCPConn(conn net.Conn) {
	c.tcpConnectionsMutex.Lock()
	delete(c.tcpConnections, conn)
	c.tcpConnectionsMutex.Unlock()
}

//AddWSConn method is adding incoming connection with login to source structure
func (c *Connections) AddWSConn(conn *websocket.Conn, userName string) {
	c.wsConnectionsMutex.Lock()
	c.wsConnections[conn] = userName
	c.wsConnectionsMutex.Unlock()
	log.Println(c.wsConnections[conn], "ADDWS")
}

//DeleteWSConn method is deleting suitable connection from wsConnections map of Connections struct
func (c *Connections) DeleteWSConnection(conn *websocket.Conn) {
	c.wsConnectionsMutex.Lock()
	delete(c.wsConnections, conn)
	c.wsConnectionsMutex.Unlock()
}

//GetUserNameByTCPConnection method returns Name of User whose connected with the TCP connection
func (c *Connections) GetUserNameByTCPConnection(conn net.Conn) string {
	c.tcpConnectionsMutex.Lock()
	defer c.tcpConnectionsMutex.Unlock()
	return c.tcpConnections[conn]
}

//GetUserNameByWSConnection method returns Name of User whose connected with the WS connection
func (c *Connections) GetUserNameByWSConnection(conn *websocket.Conn) string {
	c.wsConnectionsMutex.Lock()
	defer c.wsConnectionsMutex.Unlock()
	return c.wsConnections[conn]
}

//GetAllTCPConnections method returns slice of tcp connections
func (c *Connections) GetAllTCPConnections() map[net.Conn]string {
	return c.tcpConnections
}

//GetAllWSConnections returns slice of ws connections
func (c *Connections) GetAllWSConnections() map[*websocket.Conn]string {
	return c.wsConnections
}
