package service

import (
	"go_messenger/server/db/dbservice"
	"go_messenger/server/handlers/tcp"
	"go_messenger/server/interfaces"
	"go_messenger/server/routing"
)

//var grm interfaces.GroupMemberInterface = dbservice.GroupMember{}

//AddGroupMember add new members in spesific Group.
func AddGroupMember(userName, groupName, lastMessage string, groupMember []string, groupType uint) {
	var msg = tcp.Message{}
	var gmi interfaces.GMI = dbservice.GroupMember{}
	for _, user := range groupMember {
		ok := gmi.AddGroupMember(user, groupName, lastMessage)
		if ok {
			msg = tcp.Message{UserName: user, GroupName: groupName, Status: true}
		}
		if !ok {
			msg = tcp.Message{UserName: user, GroupName: groupName, Status: false}
		}
		routing.RoutOut(msg)
	}
}
