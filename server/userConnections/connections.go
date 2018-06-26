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
	WSConnectionsMutex  *sync.Mutex
	WSConnections       map[*websocket.Conn]string // connection:login
	TCPConnectionsMutex *sync.Mutex
	TCPConnections      map[net.Conn]string // connection:login
	OutChan             chan *serviceModels.MessageOut
}

//Function InitConnections is init for Connections struct
func InitConnections() *Connections {
	instance := Connections{}
	instance.WSConnectionsMutex = new(sync.Mutex)
	instance.WSConnections = make(map[*websocket.Conn]string, 0)
	instance.TCPConnectionsMutex = new(sync.Mutex)
	instance.TCPConnections = make(map[net.Conn]string, 0)
	instance.OutChan = make(chan *serviceModels.MessageOut, 1024)
	return &instance
}

//AddTCPConn method is adding incoming connection with login to source structure
func (c *Connections) AddTCPConn(conn net.Conn, userName string) {
	c.TCPConnectionsMutex.Lock()
	c.TCPConnections[conn] = userName
	c.TCPConnectionsMutex.Unlock()
	log.Println(c.TCPConnections[conn], "ADDTCP")
}

//DeleteTCPConn method is deleting suitable connection from TCPConnections map of Connections struct
func (c *Connections) DeleteTCPConn(conn net.Conn) {
	c.TCPConnectionsMutex.Lock()
	delete(c.TCPConnections, conn)
	c.TCPConnectionsMutex.Unlock()
}

//AddWSConn method is adding incoming connection with login to source structure
func (c *Connections) AddWSConn(conn *websocket.Conn, userName string) {
	c.WSConnectionsMutex.Lock()
	c.WSConnections[conn] = userName
	c.WSConnectionsMutex.Unlock()
	log.Println(c.WSConnections[conn], "ADDWS")
}

//DeleteWSConn method is deleting suitable connection from WSConnections map of Connections struct
func (c *Connections) DeleteWSConnection(conn *websocket.Conn) {
	c.WSConnectionsMutex.Lock()
	delete(c.WSConnections, conn)
	c.WSConnectionsMutex.Unlock()
}

//GetUserNameByTCPConnection method returns Name of User whose connected with the TCP connection
func (c *Connections) GetUserNameByTCPConnection(conn net.Conn) string {
	c.TCPConnectionsMutex.Lock()
	defer c.TCPConnectionsMutex.Unlock()
	return c.TCPConnections[conn]
}

//GetUserNameByWSConnection method returns Name of User whose connected with the WS connection
func (c *Connections) GetUserNameByWSConnection(conn *websocket.Conn) string {
	c.WSConnectionsMutex.Lock()
	defer c.WSConnectionsMutex.Unlock()
	return c.WSConnections[conn]
}

//GetAllTCPConnections method returns slice of tcp connections
func (c *Connections) GetAllTCPConnections() map[net.Conn]string {
	return c.TCPConnections
}

//GetAllWSConnections returns slice of ws connections
func (c *Connections) GetAllWSConnections() map[*websocket.Conn]string {
	return c.WSConnections
}
