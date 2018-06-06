package service

import (
<<<<<<< HEAD
=======
	"go_messenger/server/db/dbservice"
	"go_messenger/server/service/interfaces"
>>>>>>> group-chat
	"go_messenger/server/userConnections"
)

//AddGroupMember add new members in spesific Group.
func AddGroupMember(chanOut chan *userConnections.Message) {
<<<<<<< HEAD


=======
	message := <-chanOut
	var gmi interfaces.GMI = dbservice.GroupMember{}
	for _, user := range message.GroupMember {
		gmi.AddGroupMember(user, message.GroupName, message.LastMessage)
	}
>>>>>>> group-chat
}
