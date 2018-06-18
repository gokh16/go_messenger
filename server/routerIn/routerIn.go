package routerIn

import (
	"fmt"
	"go_messenger/server/service"
	"go_messenger/server/userConnections"
)

//RouterIn method which directs data to next step by action field in message structure
func RouterIn(msg *userConnections.MessageIn, chanOut chan *serviceModels.MessageOut) {

	// variable "action" is a command what to do with the structures
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
		fmt.Println("Unknown format of data from server")
	}
}
