package main

import (
	"fmt"
	"go_messenger/server/db"
	"go_messenger/server/db/dbservice"
	"go_messenger/server/handlers/tcp"
	"go_messenger/server/handlers/ws"
	"go_messenger/server/routerOut"
	"go_messenger/server/userConnections"
	"log"
	"net"
	"sync"

	"github.com/gorilla/websocket"
)

func init() {
	db.CreateDatabase()
	//db.InitDatabase()

}

func main() {
	chOut := make(chan *userConnections.Message, 1024)
	connectionList := userConnections.Connections{}
	connectionList.OutChan = chOut
	connectionList.WSConnectionsMutex = new(sync.Mutex)
	connectionList.WSConnections = make(map[*websocket.Conn]string)
	connectionList.TCPConnectionsMutex = new(sync.Mutex)
	connectionList.TCPConnections = make(map[net.Conn]string)

	routerOut.NewRouterOut(&connectionList)

	ws.NewHandlerWS(&connectionList)
	fmt.Println("good")

	tcp.NewHandlerTCP(&connectionList)
	fmt.Println("good2")

	db := dbservice.OpenConnDB()
	defer func() {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	fmt.Println("good3")

	stop := make(chan bool)
	<-stop
}

//TODO GOLINTERS
//TODO waitgroups
