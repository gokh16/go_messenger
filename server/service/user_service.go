package service

import (
	"go_messenger/server/db/dbservice"
	"go_messenger/server/models"
	"go_messenger/server/service/interfaces"
	"go_messenger/server/userConnections"
)

//CreateUser function creats a special User and makes a record in DB. It returns bool value
func CreateUser(login, password, userName, email, userIcon string, status bool, chanOut chan *userConnections.Message) {
	var ui interfaces.UI = dbservice.User{}
	message := userConnections.Message{UserName: userName, Login: login, Password: password, Email: email, UserIcon: userIcon}
	user := models.User{Login: login, Password: password, Username: userName, Email: email, Status: status, UserIcon: userIcon}
	ok := ui.CreateUser(user)
	if ok {
		message = userConnections.Message{Status: ok}
	}

	message = userConnections.Message{Status: ok}

	chanOut <- &message
}
