package service

import (
	"go_messenger/server/db/dbservice"
	"go_messenger/server/routing"
	"go_messenger/server/service/interfaces"
	"go_messenger/server/userConnections"
)

//CreateUser function creats a special User and makes a record in DB. It returns bool value
func CreateUser(login, password, userName, email, userIcon string, status bool) {
	var ui interfaces.UI = dbservice.User{}
	var message = userConnections.Message{UserName: userName, Login: login, Password: password, Email: email, UserIcon: userIcon}
	ok := ui.CreateUser(login, password, userName, email, status, userIcon)
	if ok {
		message = userConnections.Message{Status: ok}
	}

	message = userConnections.Message{Status: ok}

	routing.RouterOut(message)
}
