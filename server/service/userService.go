package service

import (
	"go_messenger/server/models"
	"go_messenger/server/service/interfaces"
	"go_messenger/server/service/serviceModels"
	"go_messenger/server/userConnections"
)

//UserService ...
type UserService struct {
	userManager    interfaces.UserManager
	groupManager   interfaces.GroupManager
	messageManager interfaces.MessageManager
}

func (u *UserService) InitUserService(ui interfaces.UserManager, gi interfaces.GroupManager, mi interfaces.MessageManager) {
	u.userManager = ui
	u.groupManager = gi
	u.messageManager = mi
}

//CreateUser function creats a special User and makes a record in DB. It returns bool value
func (u *UserService) CreateUser(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	messageOut := serviceModels.MessageOut{Action: messageIn.Action,
		Status: u.userManager.CreateUser(&messageIn.User)}
	chanOut <- &messageOut
}

//LoginUser - user's auth.
func (u *UserService) LoginUser(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	messageOut := serviceModels.MessageOut{Action: messageIn.Action}
	if messageIn.User.Password == "" || messageIn.User.Login == "" {
		messageOut.Err = "Empty Login or Password"
		messageOut.Status = false
		chanOut <- &messageOut
	}
	ok := u.userManager.LoginUser(&messageIn.User)
	if ok {
		groupList := u.groupManager.GetGroupList(&messageIn.User)
		for _, group := range groupList {
			groupOut := serviceModels.Group{GroupName: group.GroupName, GroupType: group.GroupType,
				Members:  u.groupManager.GetMemberList(&group),
				Messages: u.messageManager.GetGroupMessages(&group, messageIn.MessageLimit),
			}
			messageOut.GroupList = append(messageOut.GroupList, groupOut)
		}
		messageOut.User = u.userManager.GetUser(&messageIn.User)
		messageOut.ContactList = u.userManager.GetContactList(&messageIn.User)
	}
	messageOut.Status = ok

	chanOut <- &messageOut
}

//AddContact add spesial user to contact list of special User
func (u *UserService) AddContact(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	ok := u.userManager.AddContact(&messageIn.User, &messageIn.Contact, messageIn.RelationType)
	messageOut := serviceModels.MessageOut{User: messageIn.User,
		Status: ok, Action: messageIn.Action}
	chanOut <- &messageOut
}

//GetUsers method gets all users from DB.
func (u *UserService) GetUsers(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	messageOut := serviceModels.MessageOut{Action: messageIn.Action, User: messageIn.User}
	userList := []models.User{}
	u.userManager.GetUsers(&userList)
	messageOut.ContactList = userList
	chanOut <- &messageOut
}

//GetUser method get special user from DB.
func (u *UserService) GetUser(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	messageOut := serviceModels.MessageOut{User: messageIn.User, Action: messageIn.Action}
	messageOut.ContactList = append(messageOut.ContactList, u.userManager.GetUser(&messageIn.User))
	chanOut <- &messageOut
}

//EditUser method edit own client's user and saves it in DB.
func (u *UserService) EditUser(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	messageOut := serviceModels.MessageOut{Action: messageIn.Action}
	chanOut <- &messageOut
}

//DeleteUser method delete own account user from DB.
func (u *UserService) DeleteUser(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	messageOut := serviceModels.MessageOut{Action: messageIn.Action}
	chanOut <- &messageOut
}
