package service

import (
	"go_messenger/server/db/dbservice"
	"go_messenger/server/db/dbservice/dbInterfaces"
	"go_messenger/server/service/serviceModels"
	"go_messenger/server/userConnections"
)

//GroupService ...
type GroupService struct {
	userManager    dbInterfaces.UserManager
	groupManager   dbInterfaces.GroupManager
	messageManager dbInterfaces.MessageManager
}

//Init for GroupService struct
//func (g GroupService) initGroupService(messageIn *userConnections.MessageIn) serviceModels.MessageOut {
//	g.userManager = &dbservice.UserDBService{}
//	g.groupManager = &dbservice.GroupDBService{}
//	g.messageManager = &dbservice.MessageDBService{}
//	messageOut := serviceModels.MessageOut{Action: messageIn.Action}
//	return messageOut
//}

//CreateGroup function creats a special Group and makes a record in DB. It returns bool value
func (g GroupService) CreateGroup(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	g.userManager = dbservice.UserDBService{}
	g.groupManager = dbservice.GroupDBService{}
	g.messageManager = dbservice.MessageDBService{}
	messageOut := serviceModels.MessageOut{Action:messageIn.Action}
	switch messageIn.Group.GroupTypeID {
	// groupTypeID == 1 means privat message
	case 1:
		ok := g.groupManager.CreateGroup(&messageIn.Group)
		if ok {
			for _, member := range messageIn.Members {
				g.groupManager.AddGroupMember(&member, &messageIn.Group, &messageIn.Message)
			}
		}
		messageOut.Status = ok
		// groupType == 2 means group chat
	case 2:
		ok := g.groupManager.CreateGroup(&messageIn.Group)
		if ok {
			g.groupManager.AddGroupMember(&messageIn.Group.User, &messageIn.Group, &messageIn.Message)
		}
		messageOut.Status = ok
	}
	chanOut <- &messageOut
}

//GetGroup gets special group fo user from DB
func (g GroupService) GetGroup(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	g.groupManager = dbservice.GroupDBService{}
	g.messageManager = dbservice.MessageDBService{}
	messageOut := serviceModels.MessageOut{Action:messageIn.Action}
	groupModel := g.groupManager.GetGroup(&messageIn.Group)
	members := g.groupManager.GetMemberList(&groupModel)
	messages := g.messageManager.GetGroupMessages(&groupModel, messageIn.MessageLimit)
	groupOut := serviceModels.NewGroup(groupModel, members, messages)
	messageOut.GroupList = append(messageOut.GroupList, *groupOut)
	chanOut <- &messageOut
}

//GetGroupList gets all groups of special user from DB
func (g GroupService) GetGroupList(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	g.groupManager = dbservice.GroupDBService{}
	g.messageManager = dbservice.MessageDBService{}
	messageOut := serviceModels.MessageOut{Action:messageIn.Action}
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
func (g GroupService) AddGroupMember(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	g.groupManager = dbservice.GroupDBService{}
	for _, member := range messageIn.Members {
		g.groupManager.AddGroupMember(&member, &messageIn.Group, &messageIn.Message)
	}
}

//GetMemberList method gets all users of special group.
func (g GroupService) GetMemberList(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	g.groupManager = dbservice.GroupDBService{}
	messageOut := serviceModels.MessageOut{Recipients: g.groupManager.GetMemberList(&messageIn.Group),
		Action: messageIn.Action}
	chanOut <- &messageOut
}