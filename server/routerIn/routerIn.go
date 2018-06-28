package routerIn

import (
	"go_messenger/server/service"
	"go_messenger/server/service/interfaces"
	"go_messenger/server/service/serviceModels"
	"go_messenger/server/userConnections"
	"log"
)

var userService = service.UserService{}
var groupService = service.GroupService{}
var messageService = service.MessageService{}

//InitServices to init service.Service structure
func InitServices(ui interfaces.UserManager, gi interfaces.GroupManager, mi interfaces.MessageManager) {
	userService.InitUserService(ui, gi, mi)
	groupService.InitGroupService(ui, gi, mi)
	messageService.InitMessageService(ui, gi, mi)
}

//RouterIn is function which directs data to next step by action field in messageIn structure
func RouterIn(messageIn *userConnections.MessageIn, chanOut chan *serviceModels.MessageOut) {

	// variable "action" is a command what to do with the structures
	action := messageIn.Action

	switch action {

	case "SendMessageTo":
		go messageService.SendMessageTo(messageIn, chanOut)
	case "CreateUser":
		go userService.CreateUser(messageIn, chanOut)
	case "LoginUser":
		go userService.LoginUser(messageIn, chanOut)
	case "CreateGroup":
		go groupService.CreateGroup(messageIn, chanOut)
	case "AddGroupMember":
		go groupService.AddGroupMember(messageIn, chanOut)
	case "GetUsers":
		go userService.GetUsers(messageIn, chanOut)

	default:
		//go UnknownAction()
		log.Println("Unknown format of data from server")
	}
}
