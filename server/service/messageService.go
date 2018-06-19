package service

import (
	"go_messenger/server/db/dbservice"
	"go_messenger/server/service/serviceModels"
	"go_messenger/server/userConnections"
)

type MessageService struct {
	MessageDBService dbservice.MessageDBService
	GroupDBService   dbservice.GroupDBService
}

func (m *MessageService) SendMessageTo(msg *userConnections.MessageIn, chanOut chan *serviceModels.MessageOut) {
	//msgOut := serviceModels.MessageOut{}
	//m.MessageDBService.AddMessage(msg.Message.Content, msg.User.Username, msg.Group.GroupName, msg.Message.MessageContentType)
	//msgOut.Members = m.GroupDBService.GetGroupUserList(msg.Group.GroupName)
	//msgOut.Action = msg.Action
	//chanOut <- &msgOut
}
