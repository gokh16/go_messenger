package service

import (
	"go_messenger/server/db/dbservice"
	"go_messenger/server/service/interfaces"
	"go_messenger/server/userConnections"
)

//CreateGroup function creats a special User and makes a record in DB. It returns bool value
func CreateGroup(chanOut chan *userConnections.Message) {
	//var message = userConnections.Message{UserName: groupOwner, GroupName: groupName}
	var gmi interfaces.GMI = dbservice.GroupMember{}
	var gi interfaces.GI = dbservice.Group{}
	message := <-chanOut
	switch {
	// groupType == 0 means privat message
	case message.GroupType == 0:
		ok := gi.CreateGroup(message.GroupName, message.GroupOwner, message.GroupType)
		if ok {
			for _, user := range message.GroupMember {
				gmi.AddGroupMember(user, message.GroupName, "")
			}
			message.Status = ok
		}
		message.Status = ok
	// groupType == 1 means group chat
	case message.GroupType == 1 || message.GroupType == 2:
		ok := gi.CreateGroup(message.GroupName, message.GroupOwner, message.GroupType)
		if ok {
			gmi.AddGroupMember(message.GroupOwner, message.GroupName, "")
			message.Status = ok
		}
		message.Status = ok

	}
	chanOut <- message
}

// func GetGroupList(userName string, chanOut chan *userConnections.Message) {

// }
