package routerIn

import (
	"go_messenger/server/userConnections"
	"log"
	"go_messenger/server/service"
)

func RouterIn(c chan *userConnections.Message) {

	// variable "action" is a command what to do with the structure
	msg := <- c
	action := msg.Action

	switch action {

	case "SendMessageTo":
		go service.SendMessageTo(c)
	case "CreateUser":
		go service.CreateUser(c)
	case "CreateGroup":
		go service.CreateGroup(c)
	case "AddGroupMember":
		go service.AddGroupMember(c)

	default:
		log.Fatal("Unknown format of data")
	}
}
