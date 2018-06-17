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
	//message := <-chanOut
	user := models.User{Username: message.UserName}

	fmt.Println(message.UserName)
	fmt.Println(user.Username)
	ok := ui.CreateUser(&user)
	if ok {
		message.Status = ok
	}

	message.Status = ok
	fmt.Println("write in channel")
	chanOut <- message
}

//LoginUser ...
//todo comment here
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

//GetUsers ...
//todo comment here
func GetUsers(message *userConnections.Message, chanOut chan *userConnections.Message) {
	var ui interfaces.UI = dbservice.User{}
	users := []models.User{}
	ui.GetUsers(&users)
	for _, val := range users {
		message.GroupMember = append(message.GroupMember, val.Username)
	}
	chanOut <- message
}

//EditUser ...
//todo comment here
func EditUser() {

}

//DeleteUser ...
//todo comment here
func DeleteUser() {

}

//GetContactList ...
//todo comment here
func GetContactList() {

}

//AddContact ...
//todo comment here
func AddContact() {

}

//DeleteContact ...
//todo comment here
func DeleteContact() {

}
