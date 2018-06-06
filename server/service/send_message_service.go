package service

import (
<<<<<<< HEAD
=======
	"go_messenger/server/db/dbservice"
	"go_messenger/server/service/interfaces"
>>>>>>> group-chat
	"go_messenger/server/userConnections"
)

//SendMessageTo ...
func SendMessageTo(chanOut chan *userConnections.Message) {
<<<<<<< HEAD

=======
	message := <-chanOut
	var mi interfaces.MI = dbservice.Message{}
	var gmi interfaces.GMI = dbservice.GroupMember{}
	mi.AddMessage(message.Content, message.UserName, message.GroupName, message.ContentType)
	var groupMember = []string{}
	userList := gmi.GetGroupUserList(message.GroupName)
	for _, value := range userList {
		groupMember = append(message.GroupMember, value.Username)
	}

	message.GroupMember = groupMember
	chanOut <- message
>>>>>>> group-chat
}
