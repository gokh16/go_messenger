package main

import (
	"go_messenger/server/db"
	"go_messenger/server/db/dbservice"
	"go_messenger/server/handlers/tcp"
	"go_messenger/server/handlers/ws"
	"go_messenger/server/routerIn"
	"go_messenger/server/routerOut"
	"go_messenger/server/userConnections"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	db.CreateDatabase()
}

func main() {

	// init connections struct
	connectionList := userConnections.InitConnections()

	// init services
	routerIn.InitServices(&dbservice.UserDBService{}, &dbservice.GroupDBService{}, &dbservice.MessageDBService{})

	// init routerOut
	routerOut.InitRouterOut(connectionList)

	// init services
	routerIn.InitServices(&dbservice.UserDBService{}, &dbservice.GroupDBService{}, &dbservice.MessageDBService{})

	// run WS server
	ws.NewHandlerWS(connectionList)

	// run TCP server
	tcp.NewHandlerTCP(connectionList)

	// get DB connection
	db := dbservice.OpenConnDB()
	defer func() {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	stop := make(chan bool)
	<-stop
}
