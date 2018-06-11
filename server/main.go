package main

import (
	"fmt"
	"net"
	"sync"

	"go_messenger/server/db"
	"go_messenger/server/db/dbservice"
	"go_messenger/server/handlers/tcp"
	"go_messenger/server/handlers/ws"
	"go_messenger/server/routerOut"
	"go_messenger/server/userConnections"

	"github.com/gorilla/websocket"
)

func init() {
	db.CreateDatabase()
	db.InitDatabase()

}

func main() {
	chOut := make(chan *userConnections.Message, 1024)

	connectionList := userConnections.Connections{}

	connectionList.OutChan = chOut
	connectionList.WSConnectionsMutex = new(sync.Mutex)
	connectionList.WSConnections = make(map[*websocket.Conn]string, 0)
	connectionList.TCPConnectionsMutex = new(sync.Mutex)
	connectionList.TCPConnections = make(map[net.Conn]string, 0)

	routerOut.NewRouterOut(&connectionList)

	ws.NewWSHandler(&connectionList)
	fmt.Println("good")

	tcp.NewTCPHandler(&connectionList)
	fmt.Println("good2")

	db := dbservice.OpenConnDB()
	defer db.Close()

	fmt.Println("good3")

	stop := make(chan bool)
	<-stop
}