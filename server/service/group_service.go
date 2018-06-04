package service

import (
	"go_messenger/server/db/dbservice"
	"go_messenger/server/handlers/tcp"
	"go_messenger/server/interfaces"
	"go_messenger/server/routing"
)

//CreateGroup function creats a special User and makes a record in DB. It returns bool value
func CreateGroup(groupName, groupOwner string, groupMember []string, groupType uint) {
	var msg = tcp.Message{}
	var gmi interfaces.GMI = dbservice.GroupMember{}
	var gi interfaces.GI = dbservice.Group{}
	ok := gi.CreateGroup(groupName, groupOwner, groupType)
	if ok {
		lastMessage := ""
		for _, user := range groupMember {
			gmi.AddGroupMember(user, groupName, lastMessage)
			msg = tcp.Message{UserName: user, GroupName: groupName, Status: true}
		}
	}
	if !ok {
		msg = tcp.Message{UserName: groupOwner, GroupName: groupName, Status: false}
	}
	routing.RoutOut(msg)
}
