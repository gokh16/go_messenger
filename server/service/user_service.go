package service

import (
	"fmt"
	"go_messenger/server/db/dbservice"
	"go_messenger/server/models"
	"go_messenger/server/service/interfaces"
	"go_messenger/server/userConnections"
)

//CreateUser function creats a special User and makes a record in DB. It returns bool value
func CreateUser(message *userConnections.Message, chanOut chan *userConnections.Message) {
	fmt.Println("Service Ok")
	var ui interfaces.UI = dbservice.User{}
	user := models.User{Username: message.UserName}
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
		ui.GetContactList(message.UserName)
		gi.GetGroupList(message.UserName)

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
