package service

import (
	"fmt"
	"go_messenger/server/db/dbservice"
	"go_messenger/server/service/interfaces"
	"go_messenger/server/userConnections"
)

//CreateGroup function creats a special User and makes a record in DB. It returns bool value
func CreateGroup(message *userConnections.Message, chanOut chan *userConnections.Message) {
	fmt.Println(message.GroupName)
	var gmi interfaces.GMI = dbservice.GroupMember{}
	var gi interfaces.GI = dbservice.Group{}
	switch message.GroupType {
	// groupType == 0 means privat message
	case 1:
		ok := gi.CreateGroup(message.GroupName, message.GroupOwner, message.GroupType)
		if ok {
			for _, user := range message.GroupMember {
				gmi.AddGroupMember(user, message.GroupName, "")
			}
			message.Status = ok
		}
		message.Status = ok
	// groupType == 1 means group chat
	case 2:
		ok := gi.CreateGroup(message.GroupName, message.GroupOwner, message.GroupType)
		if ok {
			gmi.AddGroupMember(message.GroupOwner, message.GroupName, "")
			message.Status = ok
		}
		message.Status = ok

	}
	chanOut <- message
}

//GetGroup ...
//todo comment here
func GetGroup() {

}

//GetGroupList ...
//todo comment here
func GetGroupList(message *userConnections.Message, chanOut chan *userConnections.Message) {

}

//EditGroup ...
//todo comment here
func EditGroup() {

}

//DeleteGroup ...
//todo comment here
func DeleteGroup() {

}
