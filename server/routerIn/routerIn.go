package routerIn

import (
	"go_messenger/server/userConnections"
	"go_messenger/server/service"
	"fmt"
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
		fmt.Println("Unknown format of data from server")
	}
}
