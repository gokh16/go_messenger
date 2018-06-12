package main

import (
	"fmt"
	"go_messenger/server/db/dbservice"
	"go_messenger/server/handlers/tcp"
	"go_messenger/server/handlers/ws"
	"go_messenger/server/userConnections"

	"go_messenger/server/db"
	"go_messenger/server/routerOut"
)

func init() {
	db.CreateDatabase()
	db.InitDatabase()
}

func main() {

	// init connections struct
	connectionList := userConnections.InitConnections()

	// init routerOut
	routerOut.NewRouterOut(connectionList)

	// start WS server
	ws.NewWSHandler(connectionList)
	fmt.Println("good 1")

	// start TCP server
	tcp.NewTCPHandler(connectionList)
	fmt.Println("good 2")

	db := dbservice.OpenConnDB()
	defer db.Close()
	fmt.Println("good 3")

	stop := make(chan bool)
	<-stop
}
