package service

import (
	"fmt"
	"go_messenger/server/db/dbservice"
	"go_messenger/server/userConnections"
)

type GroupService struct {
	groupDBService dbservice.GroupDBService
}

func (g *GroupService) CreateGroup(msg *userConnections.Message, chanOut chan *userConnections.Message) {
	fmt.Println(msg.Group.GroupName)

	switch msg.Group.GroupType.Type {
	// groupType == 1 means privat msg
	case 1:
		ok := g.groupDBService.CreateGroup(msg.Group.GroupName, msg.Group.GroupOwnerID, msg.Group.GroupType.Type)
		if ok {
			for _, user := range msg.GroupMember.GroupMember {
				g.groupDBService.AddGroupMember(user.Username, msg.Group.GroupName, "")
			}
			msg.User.Status = ok
		}
		msg.User.Status = ok
		// groupType == 2 means group chat
	case 2:
		ok := g.groupDBService.CreateGroup(msg.Group.GroupName, msg.Group.GroupOwnerID, msg.Group.GroupType.Type)
		if ok {
			g.groupDBService.AddGroupMember(msg.Group.GroupOwnerID, msg.Group.GroupName, "")
			msg.User.Status = ok
		}
		msg.Group.GroupStatus = ok

	}
	chanOut <- msg
}

func (g *GroupService) GetGroup(message *userConnections.Message, chanOut chan *userConnections.Message) {
	// var gi interfaces.GI = dbservice.GroupDBService{}
	// var mi interfaces.MI = dbservice.MessageDBService{}
	// group := gi.GetGroup(message.GroupName)
	// groupMessages := mi.GetGroupMessages(message.GroupName)
	// groupMembers := gi.GetGroupUserList(message.GroupName)
	chanOut <- message
}

func (g *GroupService) GetGroupList(message *userConnections.Message, chanOut chan *userConnections.Message) {
	chanOut <- message
}

func (*GroupService) EditGroup(message *userConnections.Message, chanOut chan *userConnections.Message) {
	chanOut <- message
}

//AddGroupMember add new members in spesific GroupDBService.
func (g *GroupService) AddGroupMember(message *userConnections.Message, chanOut chan *userConnections.Message) {
	//var gi interfaces.GI = dbservice.GroupDBService{}
	for _, user := range message.GroupMember.GroupMember {
		g.groupDBService.AddGroupMember(user.Username, message.Group.GroupName, message.Message.LastMessage)
	}
}

