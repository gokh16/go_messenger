package routerIn

import (
	"fmt"
	"go_messenger/server/service"
	"go_messenger/server/userConnections"
)

func RouterIn(msg *userConnections.Message, chOut chan *userConnections.Message) {

	// variable "action" is a command what to do with the structure
	action := msg.Action

	switch action {

	case "SendMessageTo":
		go service.SendMessageTo(msg, chOut)
	case "CreateUser":
		go service.CreateUser(msg, chOut)
	case "CreateGroup":
		go service.CreateGroup(msg, chOut)
	case "AddGroupMember":
		go service.AddGroupMember(msg, chOut)

	default:
		fmt.Println("Unknown format of data from server")
	}
}
