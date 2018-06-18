package service

import (
	"go_messenger/server/db/dbinterfaces"
	"go_messenger/server/db/dbservice"
	"go_messenger/server/models"
	"go_messenger/server/service/interfaces"
	"go_messenger/server/service/serviceModels"
	"go_messenger/server/userConnections"
)

//CreateUser function creats a special User and makes a record in DB. It returns bool value
func CreateUser(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	var createUser dbinterfaces.UserManager = dbservice.User{}
	ok := createUser.CreateUser(&messageIn.User)
	messageOut := serviceModels.MessageOut{Status: ok, Action: messageIn.Action}
	chanOut <- &messageOut

}

func LoginUser(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	var loginUser dbinterfaces.UserManager = dbservice.User{}
	var groupList dbinterfaces.GroupManager = dbservice.Group{}
	ok := loginUser.LoginUser(&messageIn.User)
	if ok {
		var group serviceModels.Group
		groups := []serviceModels.Group{}
		user := loginUser.GetUser(&messageIn.User)
		contactList := loginUser.GetContactList(&messageIn.User)
		groups := groupList.GetGroupList(&messageIn.User)

		messageOut := serviceModels.MessageOut{User: &user, Users: contactList,
			Groups: groups, Status: ok, Action: messageIn.Action}
	}
}

func GetUsers(message *userConnections.Message, chanOut chan *userConnections.Message) {
	var ui interfaces.UI = dbservice.User{}
	user := []models.User{}
	ui.GetUsers(&user)
	for _, val := range user {
		message.GroupMember = append(message.GroupMember, val.Username)
	}
	chanOut <- message
}

// func EditUser() {

// }

// func DeleteUser() {

// }

// func GetContactList() {

// }

// func AddContact() {

// }

// func DeleteContact() {

// }
