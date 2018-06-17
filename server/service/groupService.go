package service

import (
	"fmt"
	"go_messenger/server/db/dbservice"
	"go_messenger/server/service/serviceModels"
	"go_messenger/server/userConnections"
)

type GroupService struct {
	groupDBService dbservice.GroupDBService
}

func (g *GroupService) CreateGroup(msg *userConnections.MessageIn, chanOut chan *serviceModels.MessageOut) {
	fmt.Println(msg.Group.GroupName)

	switch msg.Group.GroupType.Type {
	// groupType == 1 means privat msg
	case 1:
		ok := g.groupDBService.CreateGroup(msg.Group.GroupName, msg.Group.GroupOwnerID, msg.Group.GroupType.Type)
		if ok {
			for _, user := range msg.Member.GroupMembers {
				g.groupDBService.AddGroupMember(user.Username, msg.Group.GroupName, 0)
			}
			msg.Status = ok
		}
		msg.Status = ok
		// groupType == 2 means group chat
	case 2:
		ok := g.groupDBService.CreateGroup(msg.Group.GroupName, msg.Group.GroupOwnerID, msg.Group.GroupType.Type)
		if ok {
			g.groupDBService.AddGroupMember(msg.User.Username, msg.Group.GroupName, 0)
			msg.User.Status = ok
		}
		msg.Status = ok

	}
	chanOut <- msg
}

func (g *GroupService) GetGroup(message *userConnections.MessageIn, chanOut chan *serviceModels.MessageOut) {
	// var gi interfaces.GI = dbservice.GroupDBService{}
	// var mi interfaces.MI = dbservice.MessageDBService{}
	// group := gi.GetGroup(message.GroupName)
	// groupMessages := mi.GetGroupMessages(message.GroupName)
	// groupMembers := gi.GetGroupUserList(message.GroupName)
	chanOut <- message
}

func (g *GroupService) GetGroupList(message *userConnections.MessageIn, chanOut chan *serviceModels.MessageOut) {
	chanOut <- message
}

func (*GroupService) EditGroup(message *userConnections.MessageIn, chanOut chan *serviceModels.MessageOut) {
	chanOut <- message
}

//AddGroupMember add new members in spesific GroupDBService.
func (g *GroupService) AddGroupMember(message *userConnections.MessageIn, chanOut chan *serviceModels.MessageOutF) {
	//var gi interfaces.GI = dbservice.GroupDBService{}
	for _, user := range message.Member.GroupMembers {
		g.groupDBService.AddGroupMember(user.Username, message.Group.GroupName, message.GroupMember.LastReadMessageID)
	}
}
