package service

import (
	"go_messenger/server/db/dbservice"
	"go_messenger/server/interfaces"
	"go_messenger/server/routing"
	"go_messenger/server/userConnections"
)

//SendMessageTo ...
func SendMessageTo(content, userName, groupName string, contentType uint) {
	var mi interfaces.MI = dbservice.Message{}
	var gmi interfaces.GMI = dbservice.GroupMember{}
	mi.AddMessage(content, userName, groupName, contentType)
	var groupMember = []string{}
	userList := gmi.GetGroupUserList(groupName)
	for _, value := range userList {
		groupMember = append(groupMember, value.Username)
	}

	message := userConnections.Message{Content: content, UserName: userName, GroupName: groupName, GroupMember: groupMember}
	routing.RouterOut(message)
}
