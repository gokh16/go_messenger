package routerIn

import (
	"go_messenger/server/userConnections"
	"log"
	"go_messenger/server/service"
)

func RouterIn(msg userConnections.Message) {

	// variable "action" is a command what to do with the structure
	action := msg.Action

	switch action {

	case "SendMessageTo":
		go service.SendMessageTo(msg.Content, msg.UserName, msg.GroupName, msg.ContentType)
	case "CreateUser":
		go service.CreateUser(msg.Login, msg.Password, msg.UserName, msg.Email, msg.UserIcon, msg.Status)
	case "CreateGroup":
		go service.CreateGroup(msg.GroupName, msg.GroupOwner, msg.GroupMember, msg.GroupType)
	case "AddGroupMember":
		go service.AddGroupMember(msg.UserName, msg.GroupName, msg.LastMessage, msg.GroupMember)

	default:
		log.Fatal("Unknown format of data")
	}
}
