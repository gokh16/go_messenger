package routerIn

import (
	"go_messenger/server/service"
	"go_messenger/server/userConnections"
	"log"
)

func RouterIn(msg *userConnections.Message, chanOut chan *userConnections.Message) {

	// variable "action" is a command what to do with the structure
	action := msg.Action

	switch action {

	case "SendMessageTo":
		go service.SendMessageTo(msg, chanOut)
	case "CreateUser":
		go service.CreateUser(msg, chanOut)
	//case "LoginUser":
	//	go service.LoginUser(msg, chanOut)
	case "CreateGroup":
		go service.CreateGroup(msg, chanOut)
	//case "AddGroupMember":
	//	go service.AddGroupMember(c)
	case "GetUsers":
		go service.GetUsers(msg, chanOut)

	default:
		log.Println("Unknown format of data from server")
	}
}
