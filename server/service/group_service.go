package service

import (
	"go_messenger/server/db/dbservice"
	"go_messenger/server/interfaces"
	"go_messenger/server/routing"
	"go_messenger/server/userConnections"
)

//CreateGroup function creats a special User and makes a record in DB. It returns bool value
func CreateGroup(groupName, groupOwner string, groupMember []string, groupType uint) {
	var msg = userConnections.Message{UserName: groupOwner, GroupName: groupName}
	var gmi interfaces.GMI = dbservice.GroupMember{}
	var gi interfaces.GI = dbservice.Group{}
	switch {
	// groupType == 1 means privat message
	case groupType == 1:
		ok := gi.CreateGroup(groupName, groupOwner, groupType)
		if ok {
			//lastMessage := ""
			for _, user := range groupMember {
				gmi.AddGroupMember(user, groupName, "")
			}
			msg = userConnections.Message{Status: ok}
		}
		msg = userConnections.Message{Status: ok}
	// groupType == 2 means group chat
	case groupType == 2:
		ok := gi.CreateGroup(groupName, groupOwner, groupType)
		if ok {
			lastMessage := ""
			gmi.AddGroupMember(groupOwner, groupName, "")
			msg = userConnections.Message{Status: ok}
		}
		msg = userConnections.Message{Status: ok}

	}
	routing.RouterOut(msg)
}
