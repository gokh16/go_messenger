package service

import (
	"go_messenger/server/service/serviceModels"
	"go_messenger/server/userConnections"
	"log"
	"go_messenger/server/service/interfaces"
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
	messageOut := serviceModels.MessageOut{User: messageIn.User, Action: messageIn.Action}
	ok := g.groupManager.CreateGroup(&messageIn.Group)
	if ok {
		switch messageIn.Group.GroupTypeID {

		// groupTypeID == 1 means privat message
		case 1:
			g.messageManager.AddMessage(&messageIn.Message)
			for _, member := range messageIn.Members {
				g.groupManager.AddGroupMember(&member, &messageIn.Group, &messageIn.Message)
			}

			groupOut := serviceModels.Group{
				GroupName: messageIn.Group.GroupName,
				GroupType: messageIn.Group.GroupType,
				Members:   messageIn.Members,
				Messages:  g.messageManager.GetGroupMessages(&messageIn.Group, messageIn.MessageLimit),
			}

			messageOut.GroupList = append(messageOut.GroupList, groupOut)
			messageOut.Recipients = messageIn.Members

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
	messageOut := serviceModels.MessageOut{User: messageIn.User, Action: messageIn.Action}
	groupModel := g.groupManager.GetGroup(&messageIn.Group)
	log.Printf("GET GROUP SERVICE, group_id -> %d, group_name -> %s", groupModel.ID, groupModel.GroupName)
	members := g.groupManager.GetMemberList(&groupModel)
	messages := g.messageManager.GetGroupMessages(&groupModel, messageIn.MessageLimit)
	for i, msg := range messages {
		log.Printf("GET GROUP SERVICE MSG; %d: %s", i, msg.Content)
	}
	groupOut := serviceModels.NewGroup(groupModel, members, messages)
	messageOut.GroupList = append(messageOut.GroupList, *groupOut)
	chanOut <- &messageOut
}

//GetGroupList gets all groups of special user from DB
func (g *GroupService) GetGroupList(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	messageOut := serviceModels.MessageOut{User: messageIn.User, Action: messageIn.Action}
	groupModelList := g.groupManager.GetGroupList(&messageIn.User)
	for _, gr := range groupModelList {
		log.Printf("GET GROUP LIST SERVICE, group_id -> %d, group_name -> %s", gr.ID, gr.GroupName)
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
func (g *GroupService) GetMemberList(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	messageOut := serviceModels.MessageOut{Recipients: g.groupManager.GetMemberList(&messageIn.Group),
		Action: messageIn.Action}
	chanOut <- &messageOut
}

//EditGroup method edit owner's special group and saves changes.
func (g *GroupService) EditGroup(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	messageOut := serviceModels.MessageOut{Action: messageIn.Action}
	chanOut <- &messageOut
}

//DeleteGroup method delete owner's special group from DB.
func (g *GroupService) DeleteGroup(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	messageOut := serviceModels.MessageOut{Action: messageIn.Action}
	chanOut <- &messageOut
}
