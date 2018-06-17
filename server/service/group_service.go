package service

import (
	"go_messenger/server/db/dbservice"
	"go_messenger/server/service/interfaces"
	"go_messenger/server/userConnections"
)

//CreateGroup function creats a special Group and makes a record in DB. It returns bool value
func CreateGroup(message *userConnections.Message, chanOut chan *userConnections.Message) {
	var gi interfaces.GI = dbservice.Group{}
	switch message.GroupType {
	// groupType == 1 means privat message
	case 1:
		ok := gi.CreateGroup(message.GroupName, message.GroupOwner, message.GroupType)
		if ok {
			for _, user := range message.GroupMember {
				gi.AddGroupMember(user, message.GroupName, "")
			}
			message.Status = ok
		}
		message.Status = ok
	// groupType == 2 means group chat
	case 2:
		ok := gi.CreateGroup(message.GroupName, message.GroupOwner, message.GroupType)
		if ok {
			gi.AddGroupMember(message.GroupOwner, message.GroupName, "")
			message.Status = ok
		}
		message.Status = ok

	}
	chanOut <- message
}

func GetGroup(message *userConnections.Message, chanOut chan *userConnections.Message) {
	// var gi interfaces.GI = dbservice.Group{}
	// var mi interfaces.MI = dbservice.Message{}
	// group := gi.GetGroup(message.GroupName)
	// groupMessages := mi.GetGroupMessages(message.GroupName)
	// groupMembers := gi.GetGroupUserList(message.GroupName)
	chanOut <- message
}

func GetGroupList(message *userConnections.Message, chanOut chan *userConnections.Message) {
	chanOut <- message
}

func EditGroup(message *userConnections.Message, chanOut chan *userConnections.Message) {
	chanOut <- message
}

//AddGroupMember add new members in spesific Group.
func AddGroupMember(message *userConnections.Message, chanOut chan *userConnections.Message) {
	var gi interfaces.GI = dbservice.Group{}
	for _, user := range message.GroupMember {
		gi.AddGroupMember(user, message.GroupName, message.LastMessage)
	}
}

func GetGroupUserList() {

}

func DeleteGroupMember() {

}

func DeleteGroup(message *userConnections.Message, chanOut chan *userConnections.Message) bool {
	chanOut <- message
	return false
}
