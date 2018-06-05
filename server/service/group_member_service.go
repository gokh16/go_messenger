package service

import (
	"go_messenger/server/db/dbservice"
	"go_messenger/server/routing"
	"go_messenger/server/service/interfaces"
	"go_messenger/server/userConnections"
)

//AddGroupMember add new members in spesific Group.
func AddGroupMember(userName, groupName, lastMessage string, groupMember []string, groupType uint) {
	var msg = userConnections.Message{}
	var gmi interfaces.GMI = dbservice.GroupMember{}
	for _, user := range groupMember {
		gmi.AddGroupMember(user, groupName, lastMessage)
	}
	msg = userConnections.Message{UserName: user, GroupName: groupName, Status: true}
	routing.RouterOut(msg)

}
