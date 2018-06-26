package routerIn

import (
	"go_messenger/server/service"
	"go_messenger/server/service/serviceModels"
	"go_messenger/server/userConnections"
	"log"
)

//RouterIn is function which directs data to next step by action field in messageIn structure
func RouterIn(messageIn *userConnections.MessageIn, chanOut chan *serviceModels.MessageOut) {

	// variable "action" is a command what to do with the structures
	action := messageIn.Action

	switch action {

	case "SendMessageTo":
		go service.MessageService{}.SendMessageTo(messageIn, chanOut)
	case "CreateUser":
		go service.UserService{}.CreateUser(messageIn, chanOut)
	case "LoginUser":
		go service.UserService{}.LoginUser(messageIn, chanOut)
	case "CreateGroup":
		go service.GroupService{}.CreateGroup(messageIn, chanOut)
	case "AddGroupMember":
		go service.GroupService{}.AddGroupMember(messageIn, chanOut)
	case "GetUsers":
		go service.UserService{}.GetUsers(messageIn, chanOut)
	case "GetGroupList":
		go service.GroupService{}.GetGroupList(messageIn, chanOut)

	default:
		log.Println("Unknown format of data from server")
	}
}
