package service

import (
	"go_messenger/server/db/dbservice"
	"go_messenger/server/db/dbservice/dbInterfaces"
	"go_messenger/server/models"
	"go_messenger/server/service/serviceModels"
	"go_messenger/server/userConnections"
)

//MessageService struct of Message model on service level
type MessageService struct {
	messageManager dbInterfaces.MessageManager
	groupManager   dbInterfaces.GroupManager
}

//SendMessageTo method add message to DB and gets list of group members.
func (m MessageService) SendMessageTo(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	m.messageManager = dbservice.MessageDBService{}
	m.groupManager = dbservice.GroupDBService{}
	m.messageManager.AddMessage(&messageIn.Message)
	members := m.groupManager.GetMemberList(&messageIn.Group)
	message := []models.Message{messageIn.Message}
	groupOut := serviceModels.NewGroup(messageIn.Group, members, message)
	messageOut := serviceModels.MessageOut{User: messageIn.User,
		Recipients: members, Action: messageIn.Action, Message:messageIn.Message}
	messageOut.GroupList = append(messageOut.GroupList, *groupOut)
	chanOut <- &messageOut
}
