package routerIn

import (
	"fmt"
	"go_messenger/server/service"
	"go_messenger/server/userConnections"
	"log"
)

func RouterIn(msg *userConnections.Message, c chan *userConnections.Message) {

	// variable "action" is a command what to do with the structure
	//msg := <- c
	action := msg.Action

	switch action {

	case "SendMessageTo":
		service.SendMessageTo(msg, c)
	case "CreateUser":
		fmt.Println("RouterIn Ok")
		service.CreateUser(msg, c)
	case "CreateGroup":
		fmt.Println("RouterIn Ok")
		service.CreateGroup(msg, c)
	case "AddGroupMember":
		service.AddGroupMember(msg, c)

	default:
		log.Fatal("Unknown format of data")
	}
}
