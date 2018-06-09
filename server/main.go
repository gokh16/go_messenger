package main

import (
	"fmt"

	"go_messenger/server/db"
	"go_messenger/server/db/dbservice"
	"go_messenger/server/routerIn"
	"go_messenger/server/userConnections"
)

func init() {
	db.CreateDatabase()
	db.InitDatabase()

}

func main() {
	chOut := make(chan *userConnections.Message, 1024)
	dbservice.OpenConnDB()
	connectionList := userConnections.Connections{}
	connectionList.OutChan = chOut
	connectionList.TCPConnections = make(map[net.Conn]string, 0)
	connectionList.WSConnections = make(map[*websocket.Conn]string, 0)
	routerOut.NewRouterOut(&connectionList, chOut)
	wsStr := &ws.WSHandler{}
	go wsStr.NewWSHandler(&connectionList)
	fmt.Println("good")
	tcpStr := &tcp.TCPHandler{}
	tcpStr.NewTCPHandler(&connectionList)
	fmt.Println("good")
}
