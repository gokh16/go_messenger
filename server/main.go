package main

import (
	"fmt"
	"net"

	"go_messenger/server/handlers/tcp"
	"go_messenger/server/handlers/ws"
	"go_messenger/server/userConnections"
	"github.com/gorilla/websocket"
)

func main() {
	connectionList := userConnections.Connections{}
	connectionList.TCPConnections = make(map[net.Conn]string, 0)
	connectionList.WSConnections = make(map[*websocket.Conn]string, 0)
	wsStr := &ws.WSHandler{}
	wsStr.NewWSHandler(&connectionList)
	fmt.Println("good")
	tcpStr := &tcp.TCPHandler{}
	tcpStr.NewTCPHandler(&connectionList)
	fmt.Println("good")
}
