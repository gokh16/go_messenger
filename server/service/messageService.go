package service

import (
	"go_messenger/server/models"
	"go_messenger/server/service/interfaces"
	"go_messenger/server/service/serviceModels"
	"go_messenger/server/userConnections"
)

//MessageService struct of Message model on service level
type MessageService struct {
	userManager    interfaces.UserManager
	messageManager interfaces.MessageManager
	groupManager   interfaces.GroupManager
}

func (m *MessageService) InitMessageService(ui interfaces.UserManager, gi interfaces.GroupManager, mi interfaces.MessageManager) {
	m.userManager = ui
	m.groupManager = gi
	m.messageManager = mi
}

//SendMessageTo method add message to DB and gets list of group members.
func (m *MessageService) SendMessageTo(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	m.messageManager.AddMessage(&messageIn.Message)
	members := m.groupManager.GetMemberList(&messageIn.Group)
	message := []models.Message{messageIn.Message}
	groupOut := serviceModels.NewGroup(messageIn.Group, members, message)
	messageOut := serviceModels.MessageOut{User: messageIn.User,
		Recipients: members, Action: messageIn.Action}
	messageOut.GroupList = append(messageOut.GroupList, *groupOut)
	chanOut <- &messageOut
}