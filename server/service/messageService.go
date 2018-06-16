package service

import (
	"go_messenger/server/db/dbservice"
	"go_messenger/server/userConnections"
	"go_messenger/server/models"
)

type MessageService struct {
	MessageDBService dbservice.MessageDBService
	GroupDBService dbservice.GroupDBService
}

func (m *MessageService) SendMessageTo(msg *userConnections.Message, chanOut chan *userConnections.Message) {
	m.MessageDBService.AddMessage(msg.Message.Content, msg.User.Username, msg.Group.GroupName, msg.Message.MessageContentType)
	//groupMember := []string{}
	var groupMembers []models.User
	userList := m.GroupDBService.GetGroupUserList(msg.Group.GroupName)
	for _, user := range userList {
		groupMembers = append(msg.GroupMember.GroupMember, user)
	}
	msg.GroupMember.GroupMember = groupMembers
	chanOut <- msg
}