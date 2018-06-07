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
	user := models.User{Login: message.Login, Password: message.Password,
		Username: message.UserName, Email: message.Email,
		Status: message.Status, UserIcon: message.UserIcon}
	ok := ui.CreateUser(&user)
	if ok {
		message.Status = ok
	}

	message.Status = ok

	chanOut <- message
}

func LoginUser(chanOut chan *userConnections.Message) {
	var ui interfaces.UI = dbservice.User{}
	var gi interfaces.GI = dbservice.Group{}
	message := <-chanOut
	user := models.User{Login: message.Login, Password: message.Password}
	ok := ui.LoginUser(&user)
	if ok {
		ui.GetUser(&user)
		ui.GetContactList()

	}
}

func EditUser() {

}

func DeleteUser() {

}

func GetContactList() {

}

func AddContact() {

}

func DeleteContact() {

}
