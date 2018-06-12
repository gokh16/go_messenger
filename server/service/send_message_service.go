package service

import (
	"go_messenger/server/db/dbservice"
	"go_messenger/server/service/interfaces"
	"go_messenger/server/userConnections"
)

//SendMessageTo ...
func SendMessageTo(message *userConnections.Message, chanOut chan *userConnections.Message) {
	var mi interfaces.MI = dbservice.Message{}
	var gi interfaces.GI = dbservice.Group{}
	mi.AddMessage(message.Content, message.UserName, message.GroupName, message.ContentType)
	groupMember := []string{}
	userList := gi.GetGroupUserList(message.GroupName)
	for _, value := range userList {
		groupMember = append(message.GroupMember, value.Username)
	}
	message.GroupMember = groupMember
	chanOut <- message
}
