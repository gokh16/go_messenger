package service

import (
	"go_messenger/server/db/dbservice"
	"go_messenger/server/models"
	"go_messenger/server/service/serviceModels"
	"go_messenger/server/userConnections"
)

type MessageService struct {
	MessageDBService dbservice.MessageDBService
	GroupDBService   dbservice.GroupDBService
}

func (m *MessageService) SendMessageTo(msg *userConnections.MessageIn, chanOut chan *serviceModels.MessageOut) {
	m.MessageDBService.AddMessage(msg.Message.Content, msg.User.Username, msg.Group.GroupName, msg.Message.MessageContentType)
	var groupMembers []models.User
	userList := m.GroupDBService.GetGroupUserList(msg.Group.GroupName)
	for _, user := range userList {
		groupMembers = append(msg.Member.GroupMembers, user)
	}
	msg.Member.GroupMembers = groupMembers
	chanOut <- msg
}
