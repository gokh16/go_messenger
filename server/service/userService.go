package service

import (
	"fmt"
	"go_messenger/server/models"
	"go_messenger/server/userConnections"
	"go_messenger/server/db/dbservice"
)

type UserService struct {
	UserDBService dbservice.UserDBService
	GroupDBService dbservice.GroupDBService
}

func (u *UserService) GetUsers(msg *userConnections.Message, chanOut chan *userConnections.Message) {
	var users []models.User
	u.UserDBService.GetUsers(&users)
	for _, user := range users {
		msg.GroupMember.GroupMember = append(msg.GroupMember.GroupMember, user)
	}
	chanOut <- msg
}

func (u *UserService) CreateUser(msg *userConnections.Message, chanOut chan *userConnections.Message) {
	fmt.Println("Service Ok")
	ok := u.UserDBService.CreateUser(&msg.User)
	if ok {
		msg.User.Status = ok
	}

	msg.User.Status = ok
	fmt.Println("write in channel")
	chanOut <- msg
}

func (u *UserService) LoginUser(chanOut chan *userConnections.Message) {
	msg := <-chanOut
	ok := u.UserDBService.LoginUser(&msg.User)
	if ok {
		u.UserDBService.GetUser(&msg.User)
		u.UserDBService.GetContactList(msg.User.Username)
		u.GroupDBService.GetGroupList(msg.User.Username)
	}
}

////CreateUser function creats a special User and makes a record in DB. It returns bool value
//func (u *User) CreateUser(message *userConnections.MessageDBService, chanOut chan *userConnections.MessageDBService) {
//	fmt.Println("Service Ok")
//	ok := u.CreateUser()
//	if ok {
//		message.User.Status = ok
//	}
//
//	message.Status = ok
//	fmt.Println("write in channel")
//	chanOut <- message
//}
//
//func LoginUser(chanOut chan *userConnections.MessageDBService) {
//	var ui interfaces.UI = dbservice.User{}
//	var gi interfaces.GI = dbservice.GroupDBService{}
//	message := <-chanOut
//	user := models.User{Login: message.Login, Password: message.Password}
//	ok := ui.LoginUser(&user)
//	if ok {
//		ui.GetUser(&user)
//		ui.GetContactList(message.UserName)
//		gi.GetGroupList(message.UserName)
//
//	}
//}
//
//func GetUsers(message *userConnections.MessageDBService, chanOut chan *userConnections.MessageDBService) {
//	var ui interfaces.UI = dbservice.User{}
//	user := []models.User{}
//	ui.GetUsers(&user)
//	for _, val := range user {
//		message.GroupMember = append(message.GroupMember, val.Username)
//	}
//	chanOut <- message
//}

//func EditUser() {
//
//}
//
//func DeleteUser() {
//
//}
//
//func GetContactList() {
//
//}
//
//func AddContact() {
//
//}
//
//func DeleteContact() {
//
//}
