package service

import (
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

func (u *UserService) InitUserService (ui dbInterfaces.UserManager, mi dbInterfaces.MessageManager, gi dbInterfaces.GroupManager) {
	u.userManager = ui
	u.groupManager = gi
	u.messageManager = mi
}

//CreateUser function creats a special User and makes a record in DB. It returns bool value
func (u *UserService) CreateUser(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	ok := u.userManager.CreateUser(&messageIn.User)
	messageOut := serviceModels.MessageOut{Status: ok, Action: messageIn.Action}
	chanOut <- &messageOut
}

//LoginUser - user's auth.
func (u *UserService) LoginUser(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	messageOut := serviceModels.MessageOut{Action:messageIn.Action}
	ok := u.userManager.LoginUser(&messageIn.User)
	if ok {
		messageOut.User = *u.userManager.GetUser(&messageIn.User)
		messageOut.ContactList = u.userManager.GetContactList(&messageIn.User)
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
	ok := u.userManager.AddContact(&messageIn.User, &messageIn.Contact, messageIn.RelationType)
	messageOut := serviceModels.MessageOut{Status: ok, Action: messageIn.Action}
	chanOut <- &messageOut
}

//GetUsers method gets all users from DB.
func (u *UserService) GetUsers(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
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
	messageOut := serviceModels.MessageOut{Action: messageIn.Action}
	messageOut.ContactList = append(messageOut.ContactList, *u.userManager.GetUser(&messageIn.User))
	chanOut <- &messageOut
}