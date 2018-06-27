package routerIn

import (
	"go_messenger/server/service"
	"go_messenger/server/service/serviceModels"
	"go_messenger/server/userConnections"
	"log"
	"go_messenger/server/db/dbservice/dbInterfaces"
)

var userService = service.UserService{}
var messageService = service.MessageService{}
var groupService = service.GroupService{}

func InitServices(ui dbInterfaces.UserManager, mi dbInterfaces.MessageManager, gi dbInterfaces.GroupManager) {
	userService.InitUserService(ui, mi, gi)
	messageService.InitMessageService(mi, gi)
	groupService.InitGroupService(ui, mi, gi)
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
	case "GetGroupList":
		go groupService.GetGroupList(messageIn, chanOut)

	default:
		log.Println("Unknown format of data from server")
	}
}
