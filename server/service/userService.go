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
	messageOut := serviceModels.MessageOut{Action: messageIn.Action}
	if messageIn.User.Password == "" || messageIn.User.Login == "" {
		messageOut.Err = "Emty Login or Password"
		messageOut.Status = false
		chanOut <- &messageOut
	}
	ok, err := u.userManager.CreateUser(&messageIn.User)
	if err != nil {
		messageOut.Err = "DBError when CreateUser. " + err.Error()
	}
	messageOut.Status = ok
	chanOut <- &messageOut
}

//LoginUser - user's auth.
func (u *UserService) LoginUser(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	messageOut := serviceModels.MessageOut{Action: messageIn.Action}
	if messageIn.User.Password == "" || messageIn.User.Login == "" {
		messageOut.Err = "Emty Login or Password"
		messageOut.Status = false
		chanOut <- &messageOut
		return
	}

	ok, err := u.userManager.LoginUser(&messageIn.User)
	if err != nil {
		messageOut.Err = "Error when LoginUser. " + err.Error()
	}

	if ok {
		groupList := u.groupManager.GetGroupList(&messageIn.User)

		for _, group := range groupList {
			groupOut := serviceModels.Group{GroupName: group.GroupName, GroupType: group.GroupType,
				Members:  u.groupManager.GetMemberList(&group),
				Messages: u.messageManager.GetGroupMessages(&group, messageIn.MessageLimit),
			}
			messageOut.GroupList = append(messageOut.GroupList, groupOut)
		}

		messageOut.User, err = u.userManager.GetAccount(&messageIn.User)
		if err != nil {
			messageOut.Err = "Error when GetAccount. " + err.Error()
		}

		messageOut.ContactList, err = u.userManager.GetContactList(&messageIn.User)
		if err != nil {
			messageOut.Err = "Error when GetContactList. " + err.Error()
		}
	}

	messageOut.User = messageIn.User
	messageOut.Status = ok

	chanOut <- &messageOut
}

//AddContact add spesial user to contact list of special User
func (u *UserService) AddContact(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	messageOut := serviceModels.MessageOut{User: messageIn.User,
		Action: messageIn.Action}
	ok, err := u.userManager.AddContact(&messageIn.User, &messageIn.Contact, messageIn.RelationType)
	if err != nil {
		messageOut.Err = "Error when AddContact" + err.Error()
	}
	messageOut.Status = ok
	chanOut <- &messageOut
}

func (u *UserService) GetContactList(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	messageOut := serviceModels.MessageOut{User: messageIn.User,
		Action: messageIn.Action}
	var err error
	messageOut.ContactList, err = u.userManager.GetContactList(&messageIn.User)
	if err != nil {
		messageOut.Err = "Error when GetContactList" + err.Error()
	}
	chanOut <- &messageOut
}

func (u *UserService) DeleteContact(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	messageOut := serviceModels.MessageOut{
		Action: messageIn.Action,
	}
	var err error
	messageOut.Status, err = u.userManager.DeleteContact(&messageIn.User, &messageIn.Contact)
	if err != nil {
		messageOut.Err = err.Error()
	}
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
// func (u *UserService) GetUser(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
// 	messageOut := serviceModels.MessageOut{User: messageIn.User, Action: messageIn.Action}
// 	messageOut.ContactList = append(messageOut.ContactList, u.userManager.GetUser(&messageIn.User))
// 	chanOut <- &messageOut
// }

//DeleteUser method delete own account from DB.
func (u *UserService) DeleteUser(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	messageOut := serviceModels.MessageOut{
		Action: messageIn.Action,
	}
	var err error
	messageOut.Status, err = u.userManager.DeleteUser(&messageIn.User)
	if err != nil {
		messageOut.Err = err.Error()
	}
	chanOut <- &messageOut
}
