package service

import (
	"go_messenger/server/db/dbservice"
	"go_messenger/server/db/dbservice/dbInterfaces"
	"go_messenger/server/models"
	"go_messenger/server/service/serviceModels"
	"go_messenger/server/userConnections"
)

//UserService ...
type UserService struct {
	userManager    dbInterfaces.UserManager
	groupManager   dbInterfaces.GroupManager
	messageManager dbInterfaces.MessageManager
}

//CreateUser function creats a special User and makes a record in DB. It returns bool value
func (u *UserService) CreateUser(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	u.userManager = dbservice.UserDBService{}
	ok := u.userManager.CreateUser(&messageIn.User)
	messageOut := serviceModels.MessageOut{Status: ok, Action: messageIn.Action}
	chanOut <- &messageOut
}

//LoginUser - user's auth.
func (u *UserService) LoginUser(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	u.userManager = dbservice.UserDBService{}
	u.groupManager = dbservice.GroupDBService{}
	u.messageManager = dbservice.MessageDBService{}
	messageOut := serviceModels.MessageOut{Action: messageIn.Action}
	ok := u.userManager.LoginUser(&messageIn.User)
	if ok {
		messageOut = serviceModels.MessageOut{User: *u.userManager.GetUser(&messageIn.User),
			ContactList: u.userManager.GetContactList(&messageIn.User)}
		groupList := u.groupManager.GetGroupList(&messageIn.User)
		for _, group := range groupList {
			groupOut := serviceModels.Group{GroupName: group.GroupName, GroupType: group.GroupType,
				Members:  u.groupManager.GetMemberList(&group),
				Messages: u.messageManager.GetGroupMessages(&group, messageIn.MessageLimit),
			}
			messageOut.GroupList = append(messageOut.GroupList, groupOut)
		}
	}
	messageOut.Status = ok
	chanOut <- &messageOut
}

//AddContact add spesial user to contact list of special User
func (u *UserService) AddContact(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	u.userManager = dbservice.UserDBService{}
	ok := u.userManager.AddContact(&messageIn.User, &messageIn.Contact, messageIn.RelationType)
	messageOut := serviceModels.MessageOut{Status: ok, Action: messageIn.Action}
	chanOut <- &messageOut
}

//GetUsers method gets all users from DB.
func (u *UserService) GetUsers(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	u.userManager = dbservice.UserDBService{}
	userList := []models.User{}
	u.userManager.GetUsers(&userList)
	messageOut := serviceModels.MessageOut{Action: messageIn.Action}
	for _, user := range userList {
		messageOut.ContactList = append(messageOut.ContactList, user)
	}
	chanOut <- &messageOut
}

//GetUser method get special user from DB.
func (u *UserService) GetUser(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	u.userManager = dbservice.UserDBService{}
	messageOut := serviceModels.MessageOut{Action: messageIn.Action}
	messageOut.ContactList = append(messageOut.ContactList, *u.userManager.GetUser(&messageIn.User))
	chanOut <- &messageOut
}

//GetContactList gets contact list of special user from DB.
func (u *UserService) GetContactList(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	u.userManager = dbservice.UserDBService{}
	messageOut := serviceModels.MessageOut{Action: messageIn.Action}
	messageOut.ContactList = u.userManager.GetContactList(&messageIn.User)
	chanOut <- &messageOut
}
