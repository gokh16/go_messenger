package service

import (
	"errors"
	"go_messenger/server/service/interfaces"
	"go_messenger/server/service/serviceModels"
	"go_messenger/server/userConnections"
	"log"
)

//GroupService ...
type GroupService struct {
	userManager    interfaces.UserManager
	groupManager   interfaces.GroupManager
	messageManager interfaces.MessageManager
}

//InitGroupService ...
func (g *GroupService) InitGroupService(ui interfaces.UserManager, gi interfaces.GroupManager, mi interfaces.MessageManager) {
	g.userManager = ui
	g.groupManager = gi
	g.messageManager = mi
}

//CreateGroup function creats a special Group and makes a record in DB. It returns bool value
func (g *GroupService) CreateGroup(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	messageOut := serviceModels.MessageOut{User: messageIn.User, Action: messageIn.Action}
	ok, err := g.groupManager.CreateGroup(&messageIn.Group)
	if err != nil {
		var serviceErr = ErrorService{}
		custErr := errors.New("Can't create group")
		serviceErr.SendError(custErr, messageIn.User, chanOut)
		return
	}
	if ok {
		for _, member := range messageIn.Members {
			g.groupManager.AddGroupMember(&member, &messageIn.Group, &messageIn.Message)

		}
	}
	messageOut.Status = ok
	chanOut <- &messageOut
}

//GetGroup gets special group fo user from DB
func (g *GroupService) GetGroup(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	messageOut := serviceModels.MessageOut{User: messageIn.User, Action: messageIn.Action}
	groupModel, err := g.groupManager.GetGroup(&messageIn.Group)
	if err != nil {
		var serviceErr = ErrorService{}
		custErr := errors.New("Can't get group")
		serviceErr.SendError(custErr, messageIn.User, chanOut)
		return
	}
	log.Printf("GET GROUP SERVICE, group_id -> %d, group_name -> %s", groupModel.ID, groupModel.GroupName)
	members, err := g.groupManager.GetMemberList(&groupModel)
	if err != nil {
		var serviceErr = ErrorService{}
		custErr := errors.New("Can't get member list")
		serviceErr.SendError(custErr, messageIn.User, chanOut)
		return
	}
	messages := g.messageManager.GetGroupMessages(&groupModel, messageIn.MessageLimit)
	groupOut := serviceModels.NewGroup(groupModel, members, messages)
	messageOut.GroupList = append(messageOut.GroupList, *groupOut)
	chanOut <- &messageOut
}

//GetGroupList gets all groups of special user from DB
func (g *GroupService) GetGroupList(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	messageOut := serviceModels.MessageOut{User: messageIn.User, Action: messageIn.Action}
	groupModelList, err := g.groupManager.GetGroupList(&messageIn.User)
	if err != nil {
		var serviceErr = ErrorService{}
		custErr := errors.New("Can't get group list")
		serviceErr.SendError(custErr, messageIn.User, chanOut)
		return
	}
	for _, group := range groupModelList {
		log.Printf("GET GROUP LIST SERVICE, group_id -> %d, group_name -> %s", group.ID, group.GroupName)
		members, err := g.groupManager.GetMemberList(&group)
		if err != nil {
			var serviceErr = ErrorService{}
			custErr := errors.New("Can't get member list")
			serviceErr.SendError(custErr, messageIn.User, chanOut)
			return
		}
		messages := g.messageManager.GetGroupMessages(&group, messageIn.MessageLimit)
		groupOut := serviceModels.NewGroup(group, members, messages)
		messageOut.GroupList = append(messageOut.GroupList, *groupOut)
	}
	messageOut.Status = true
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
	messageOut := serviceModels.MessageOut{User: messageIn.User,
		Action: messageIn.Action}
	recipients, err := g.groupManager.GetMemberList(&messageIn.Group)
	if err != nil {
		var serviceErr = ErrorService{}
		custErr := errors.New("Can't get member list")
		serviceErr.SendError(custErr, messageIn.User, chanOut)
		return
	}
	messageOut.Recipients = recipients

	chanOut <- &messageOut
}

//EditGroup method edit owner's special group and saves changes.
func (g *GroupService) EditGroup(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	ok := g.groupManager.EditGroup(&messageIn.Group)
	messageOut := serviceModels.MessageOut{Action: messageIn.Action, User: messageIn.User, Status: ok}
	chanOut <- &messageOut
}

//DeleteGroup method delete owner's special group from DB.
func (g *GroupService) DeleteGroup(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	messageOut := serviceModels.MessageOut{Action: messageIn.Action}
	chanOut <- &messageOut
}
