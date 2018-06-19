package service

import (
	"go_messenger/server/db/dbservice"
	"go_messenger/server/db/dbservice/dbInterfaces"
	"go_messenger/server/service/serviceModels"
	"go_messenger/server/userConnections"
)

//MessageService struct of Message model on service level
type MessageService struct {
	messageManager dbInterfaces.MessageManager
	groupManager   dbInterfaces.GroupManager
}

//SendMessageTo method add message to DB and gets list of group members.
func (s MessageService) SendMessageTo(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	s.messageManager = dbservice.Message{}
	s.groupManager = dbservice.Group{}
	s.messageManager.AddMessage(&messageIn.Message)
	groupOut := serviceModels.Group{}
	groupOut.Members = s.groupManager.GetMemberList(&messageIn.Group)
	messageOut := serviceModels.MessageOut{User: messageIn.User,
		Members: groupOut.Members, Action: messageIn.Action}
	messageOut.GroupList = append(messageOut.GroupList, groupOut)
	chanOut <- &messageOut
}
