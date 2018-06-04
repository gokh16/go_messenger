package service

import (
	"go_messenger/server/db/dbservice"
	"go_messenger/server/handlers/tcp"
	"go_messenger/server/interfaces"
	"go_messenger/server/routing"
)

var message = tcp.Message{}

//CreateUser function creats a special User and makes a record in DB. It returns bool value
func CreateUser(login, password, userName, email, userIcon string, status bool) {
	var ui interfaces.UI = dbservice.User{}
	ok := ui.CreateUser(login, password, userName, email, status, userIcon)
	if ok {
		message = tcp.Message{UserName: userName, Status: true}
	}
	if !ok {
		message = tcp.Message{UserName: userName, Status: false}
	}

	routing.RoutOut(message)
}
