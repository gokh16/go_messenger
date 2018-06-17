package userConnections

import (
	"fmt"
	"net"
	"sync"

	"github.com/gorilla/websocket"
)

//Connections is a structure with connections and channel for write out data
type Connections struct {
	WSConnectionsMutex  *sync.Mutex
	WSConnections       map[*websocket.Conn]string // connection:login
	TCPConnectionsMutex *sync.Mutex
	TCPConnections      map[net.Conn]string // connection:login
	OutChan             chan *Message
}

//AddTCPConn method is adding incoming connection with login to source structure
func (c *Connections) AddTCPConn(conn net.Conn, userName string) {
	c.TCPConnectionsMutex.Lock()
	c.TCPConnections[conn] = userName
	c.TCPConnectionsMutex.Unlock()
	fmt.Println(c.TCPConnections, "ADDTCP")
}

//AddWSConn method is adding incoming connection with login to source structure
func (c *Connections) AddWSConn(conn *websocket.Conn, userName string) {
	c.WSConnectionsMutex.Lock()
	defer c.WSConnectionsMutex.Unlock()
	c.WSConnections[conn] = userName
}

//GetAllTCPConnections method returns slice of tcp connections
func (c *Connections) GetAllTCPConnections() map[net.Conn]string {
	c.TCPConnectionsMutex.Lock()
	defer c.TCPConnectionsMutex.Unlock()
	return c.TCPConnections
}

//GetAllWSConnections returns slice of ws connections
func (c *Connections) GetAllWSConnections() map[*websocket.Conn]string {
	c.WSConnectionsMutex.Lock()
	defer c.WSConnectionsMutex.Unlock()
	return c.WSConnections
}
