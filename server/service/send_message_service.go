package service

import (
	"go_messenger/server/userConnections"
	"go_messenger/server/service/interfaces"
	"go_messenger/server/db/dbservice"
)

//SendMessageTo ...
func SendMessageTo(chanOut chan *userConnections.Message) {
	message := <-chanOut
	var mi interfaces.MI = dbservice.Message{}
	var gmi interfaces.GMI = dbservice.GroupMember{}
	mi.AddMessage(message.Content, message.UserName, message.GroupName, message.ContentType)
	var groupMember = []string{}
	userList := gmi.GetGroupUserList(message.GroupName)
	for _, value := range userList {
		groupMember = append(message.GroupMember, value.Username)
	}

	message.GroupMember = groupMember
	chanOut <- message
}
