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
	messageOut := serviceModels.MessageOut{User: messageIn.User,
		Action: messageIn.Action, Message: messageIn.Message}
	m.messageManager.AddMessage(&messageIn.Message)
	members, err := m.groupManager.GetMemberList(&messageIn.Group)
	if err != nil {
		messageOut.Err = err.Error()
		chanOut <- &messageOut
		return
	}
	message := []models.Message{messageIn.Message}
	group, err := m.groupManager.GetGroup(&messageIn.Group)
	if err != nil {
		messageOut.Err = err.Error()
		chanOut <- &messageOut
		return
	}
	groupOut := serviceModels.NewGroup(group, members, message)
	messageOut.Recipients = members
	messageOut.GroupList = append(messageOut.GroupList, *groupOut)
	chanOut <- &messageOut
}
