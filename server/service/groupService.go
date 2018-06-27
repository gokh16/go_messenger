package service

import (
	"go_messenger/server/service/interfaces"
	"go_messenger/server/service/serviceModels"
	"go_messenger/server/userConnections"
)

//GroupService ...
type GroupService struct {
	userManager    interfaces.UserManager
	groupManager   interfaces.GroupManager
	messageManager interfaces.MessageManager
}

func (g *GroupService) InitGroupService(ui interfaces.UserManager, gi interfaces.GroupManager, mi interfaces.MessageManager) {
	g.userManager = ui
	g.groupManager = gi
	g.messageManager = mi
}

//CreateGroup function creats a special Group and makes a record in DB. It returns bool value
func (g *GroupService) CreateGroup(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	messageOut := serviceModels.MessageOut{Action: messageIn.Action}
	ok := g.groupManager.CreateGroup(&messageIn.Group)
	if ok {
		switch messageIn.Group.GroupTypeID {
		// groupTypeID == 1 means privat message
		case 1:
			for _, member := range messageIn.Members {
				g.groupManager.AddGroupMember(&member, &messageIn.Group, &messageIn.Message)
			}
			// groupType == 2 means group chat
		case 2:
			g.groupManager.AddGroupMember(&messageIn.Group.User, &messageIn.Group, &messageIn.Message)
		}
	}
	messageOut.Status = ok
	chanOut <- &messageOut
}

//GetGroup gets special group fo user from DB
func (g *GroupService) GetGroup(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	messageOut := serviceModels.MessageOut{Action: messageIn.Action}
	groupModel := g.groupManager.GetGroup(&messageIn.Group)
	members := g.groupManager.GetMemberList(&groupModel)
	messages := g.messageManager.GetGroupMessages(&groupModel, messageIn.MessageLimit)
	groupOut := serviceModels.NewGroup(groupModel, members, messages)
	messageOut.GroupList = append(messageOut.GroupList, *groupOut)
	chanOut <- &messageOut
}

//GetGroupList gets all groups of special user from DB
func (g *GroupService) GetGroupList(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	messageOut := serviceModels.MessageOut{User: messageIn.User, Action: messageIn.Action}
	groupModelList := g.groupManager.GetGroupList(&messageIn.User)
	for _, gr := range groupModelList {
		members := g.groupManager.GetMemberList(&gr)
		messages := g.messageManager.GetGroupMessages(&gr, messageIn.MessageLimit)
		groupOut := serviceModels.NewGroup(gr, members, messages)
		messageOut.GroupList = append(messageOut.GroupList, *groupOut)
	}
	chanOut <- &messageOut
}

//AddGroupMember add new members in spesific Group.
func (g *GroupService) AddGroupMember(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	for _, member := range messageIn.Members {
		g.groupManager.AddGroupMember(&member, &messageIn.Group, &messageIn.Message)
	}
}

//GetMemberList method gets all users of special group.
func (s *GroupService) GetMemberList(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	messageOut := serviceModels.MessageOut{Recipients: s.groupManager.GetMemberList(&messageIn.Group),
		Action: messageIn.Action}
	chanOut <- &messageOut
}

//EditGroup method edit owner's special group and saves changes.
func (s *GroupService) EditGroup(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	messageOut := serviceModels.MessageOut{Action: messageIn.Action}
	chanOut <- &messageOut
}

//DeleteGroup method delete owner's special group from DB.
func (s *GroupService) DeleteGroup(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	messageOut := serviceModels.MessageOut{Action: messageIn.Action}
	chanOut <- &messageOut
}
