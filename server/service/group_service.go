package service

import (
	"go_messenger/server/db/dbservice"
	"go_messenger/server/service/interfaces"
	"go_messenger/server/userConnections"
)

//CreateGroup function creats a special User and makes a record in DB. It returns bool value
func CreateGroup(groupName, groupOwner string, groupMember []string, groupType uint, chanOut chan *userConnections.Message) {
	var message = userConnections.Message{UserName: groupOwner, GroupName: groupName}
	var gmi interfaces.GMI = dbservice.GroupMember{}
	var gi interfaces.GI = dbservice.Group{}
	switch {
	// groupType == 0 means privat message
	case groupType == 0:
		ok := gi.CreateGroup(groupName, groupOwner, groupType)
		if ok {
			//lastMessage := ""
			for _, user := range groupMember {
				gmi.AddGroupMember(user, groupName, "")
			}
			message = userConnections.Message{Status: ok}
		}
		message = userConnections.Message{Status: ok}
	// groupType == 1 means group chat
	case groupType == 1 || groupType == 2:
		ok := gi.CreateGroup(groupName, groupOwner, groupType)
		if ok {
			gmi.AddGroupMember(groupOwner, groupName, "")
			message = userConnections.Message{Status: ok}
		}
		message = userConnections.Message{Status: ok}

	}
	chanOut <- &message
}
