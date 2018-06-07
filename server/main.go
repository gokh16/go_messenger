package main

import (
	"fmt"
	"net"
	"go_messenger/server/userConnections"
	"github.com/gorilla/websocket"
	"go_messenger/server/handlers/tcp"
	"go_messenger/server/handlers/ws"
	"go_messenger/server/routerOut"
)

func main() {
	connectionList := userConnections.Connections{}
	connectionList.OutChan = make(chan *userConnections.Message, 1024)
	connectionList.TCPConnections = make(map[net.Conn]string, 0)
	connectionList.WSConnections = make(map[*websocket.Conn]string, 0)
	wsStr := &ws.WSHandler{}
	go wsStr.NewWSHandler(&connectionList)
	fmt.Println("good")
	tcpStr := &tcp.TCPHandler{}
	tcpStr.NewTCPHandler(&connectionList)
	fmt.Println("good")
	go routerOut.NewRouterOut(&connectionList)
}
