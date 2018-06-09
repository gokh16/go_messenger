package service

import (
	"go_messenger/server/db/dbservice"
	"go_messenger/server/service/interfaces"
	"go_messenger/server/userConnections"
)

//AddGroupMember add new members in spesific Group.
func AddGroupMember(message *userConnections.Message, chanOut chan *userConnections.Message) {
	var gmi interfaces.GMI = dbservice.GroupMember{}
	for _, user := range message.GroupMember {
		gmi.AddGroupMember(user, message.GroupName, message.LastMessage)
	}
}

func GetGroupUserList() {

}

func DeleteGroupMember() {

}
