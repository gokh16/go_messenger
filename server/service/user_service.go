package service

import (
	"go_messenger/server/db/dbservice"
	"go_messenger/server/models"
	"go_messenger/server/service/interfaces"
	"go_messenger/server/userConnections"
)

//CreateUser function creats a special User and makes a record in DB. It returns bool value
func CreateUser(chanOut chan *userConnections.Message) {
	var ui interfaces.UI = dbservice.User{}
	message := <-chanOut
	//message := userConnections.Message{UserName: userName, Login: login, Password: password, Email: email, UserIcon: userIcon}
	user := models.User{Login: message.Login, Password: message.Password, Username: message.UserName, Email: message.Email,
		Status: message.Status, UserIcon: message.UserIcon}
	ok := ui.CreateUser(&user)
	if ok {
		message.Status = ok
	}

	message.Status = ok

	chanOut <- message
}
