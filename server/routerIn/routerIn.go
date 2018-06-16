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
		go service.MessageService{}.SendMessageTo(msg, chanOut)
	case "CreateUser":
		go service.UserService{}.CreateUser(msg, chanOut)
	case "LoginUser":
		go service.UserService{}.LoginUser(chanOut)
	case "CreateGroup":
		go service.GroupService{}.CreateGroup(msg, chanOut)
	case "AddGroupMember":
		go service.GroupService{}.AddGroupMember(msg, chanOut)
	case "GetUsers":
		go service.UserService{}.GetUsers(msg, chanOut)

	default:
		log.Println("Unknown format of data from server")
	}
}
