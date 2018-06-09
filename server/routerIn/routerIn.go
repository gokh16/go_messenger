package routerIn

import (
	"fmt"
	"go_messenger/server/service"
	"go_messenger/server/userConnections"
)

func RouterIn(msg *userConnections.Message, c chan *userConnections.Message) {

	// variable "action" is a command what to do with the structure
	//msg := <- c
	action := msg.Action

	switch action {

	case "SendMessageTo":
		go service.SendMessageTo(msg, str)
	//case "CreateUser":
	//	go service.CreateUser(c)
	//case "CreateGroup":
	//	go service.CreateGroup(c)
	//case "AddGroupMember":
	//	go service.AddGroupMember(c)

	default:
		fmt.Println("Unknown format of data from server")
	}
}
