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
		return
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
		messageOut.Err = "Empty Login or Password"
		messageOut.Status = false
		chanOut <- &messageOut
		return
	}

	ok, err := u.userManager.LoginUser(&messageIn.User)
	if err != nil {
		messageOut.Err = "Error when LoginUser. " + err.Error()
		chanOut <- &messageOut
		return
	}

	if ok {
		groupList, err := u.groupManager.GetGroupList(&messageIn.User)
		if err != nil {
			messageOut.Err = err.Error()
			chanOut <- &messageOut
			return
		}

		for _, group := range groupList {
			members, err := u.groupManager.GetMemberList(&group)
			if err != nil {
				messageOut.Err = err.Error()
				chanOut <- &messageOut
				return
			}
			groupOut := serviceModels.Group{GroupName: group.GroupName, GroupType: group.GroupType,
				Members:  members,
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
		messageOut.Err = "Error when DeleteContact" + err.Error()
	}
	chanOut <- &messageOut
}

func (u *UserService) GetContactList(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	messageOut := serviceModels.MessageOut{User: messageIn.User,
		Action: messageIn.Action}
	messageOut.ContactList = u.userManager.GetContactList(&messageIn.User)
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
	user, err := u.userManager.GetUser(&messageIn.User)
	if err != nil {
		messageOut.Err = err.Error()
		chanOut <- &messageOut
		return
	}
	messageOut.ContactList = append(messageOut.ContactList, user)
	chanOut <- &messageOut
}

//EditUser method edit own client's user and saves it in DB.
func (u *UserService) EditUser(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	updatedUser := u.userManager.EditUser(&messageIn.User)
	messageOut := serviceModels.MessageOut{Action: messageIn.Action, User: messageIn.User, Status: updatedUser.Status}
	chanOut <- &messageOut
}

//DeleteUser method delete own account from DB.
func (u *UserService) DeleteUser(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	messageOut := serviceModels.MessageOut{
		Action: messageIn.Action,
	}
	var err error
	messageOut.Status, err = u.userManager.DeleteUser(&messageIn.User)
	if err != nil {
		messageOut.Err = "Error when DeleteUser" + err.Error()
	}
	chanOut <- &messageOut
}
