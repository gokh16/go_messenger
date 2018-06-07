package routerIn

import (
	"go_messenger/server/userConnections"
	"fmt"
)

func RouterIn(c chan *userConnections.Message) {

	// variable "action" is a command what to do with the structure
	msg := <- c
	fmt.Println(msg.Content)
	//action := msg.Action
	//switch action {
	//
	//case "SendMessageTo":
	//	go service.SendMessageTo(c)
	//case "CreateUser":
	//	go service.CreateUser(c)
	//case "CreateGroup":
	//	go service.CreateGroup(c)
	//case "AddGroupMember":
	//	go service.AddGroupMember(c)
	//
	//default:
	//	log.Fatal("Unknown format of data")
	//}
}
