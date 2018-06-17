package service

import (
	"go_messenger/server/db/dbservice"
	"go_messenger/server/service/interfaces"
	"go_messenger/server/userConnections"
)

//SendMessageTo ...
func SendMessageTo(message *userConnections.Message, chanOut chan *userConnections.Message) {
	var mi interfaces.MI = dbservice.Message{}
	var gmi interfaces.GMI = dbservice.GroupMember{}
	mi.AddMessage(message.Content, message.UserName, message.GroupName, message.ContentType)
	groupMember := []string{}
	userList := gmi.GetGroupUserList(message.GroupName)
	for _, value := range userList {
		groupMember = append(groupMember, value.Username)
	}
	message.GroupMember = groupMember
	chanOut <- message
}
