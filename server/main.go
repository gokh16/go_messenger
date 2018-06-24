package main

import (
	"fmt"
	"go_messenger/server/handlers/tcp"
	"go_messenger/server/handlers/ws"
	"go_messenger/server/routerOut"
	"go_messenger/server/userConnections"
	"go_messenger/server/db"
	"go_messenger/server/db/dbservice"
	"log"
)

func init() {
	db.CreateDatabase()
	//db.InitDatabase()

}

func main() {

	// init connections struct
	connectionList := userConnections.InitConnections()

	// init routerOut
	routerOut.InitRouterOut(connectionList)

	ws.NewHandlerWS(connectionList)
	fmt.Println("WS started : Ok!")

	tcp.NewHandlerTCP(connectionList)
	fmt.Println("TCP started : Ok!")

	db := dbservice.OpenConnDB()
	defer func() {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	fmt.Println("DB opened : Ok!")

	stop := make(chan bool)
	<-stop
}