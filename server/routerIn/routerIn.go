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
		go service.SendMessageTo(msg, c)
	case "CreateUser":
<<<<<<< HEAD
		go service.CreateUser(c)
	case "CreateGroup":
		go service.CreateGroup(c)
	case "AddGroupMember":
		go service.AddGroupMember(c)
=======
		go service.CreateUser(msg, c)
	//case "LoginUser":
	//	go service.LoginUser(msg, c)
	case "CreateGroup":
		go service.CreateGroup(msg, c)
	//case "AddGroupMember":
	//	go service.AddGroupMember(c)
	case "GetUsers":
		go service.GetUsers(msg, c)
>>>>>>> 2796992922c38b25ca0fdfcad6d75f378c19f7c5

	default:
		fmt.Println("Unknown format of data from server")
	}
}
