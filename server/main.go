package main

import (
	"go_messenger/server/handlers/tcp"
	"go_messenger/server/handlers/ws"
	"go_messenger/server/userConnections"

	"go_messenger/server/routerOut"
	"log"
	"go_messenger/server/db/dbservice"
	"go_messenger/server/db"
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

	// start WS server
	ws.NewWSHandler(connectionList)
	log.Println("WS handler Ok! : main")

	// start TCP server
	tcp.NewTCPHandler(connectionList)
	log.Println("TCP handler Ok! : main")

	db := dbservice.OpenConnDB()
	defer db.Close()
	log.Println("DB connection Ok! : main")

	stop := make(chan bool)
	<-stop
}
