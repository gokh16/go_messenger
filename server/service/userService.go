package service

import (
	"fmt"
	"go_messenger/server/db/dbservice"
	"go_messenger/server/models"
	"go_messenger/server/service/serviceModels"
	"go_messenger/server/userConnections"
)

type UserService struct {
	UserDBService  dbservice.UserDBService
	GroupDBService dbservice.GroupDBService
}

func (u *UserService) GetUsers(msg *userConnections.MessageIn, chanOut chan *serviceModels.MessageOut) {
	var users []models.User
	u.UserDBService.GetUsers(&users)
	for _, user := range users {
		msg.Member.GroupMembers = append(msg.Member.GroupMembers, user)
	}
	chanOut <- msg
}

func (u *UserService) CreateUser(msg *userConnections.MessageIn, chanOut chan *serviceModels.MessageOut) {
	fmt.Println("Service Ok")
	ok := u.UserDBService.CreateUser(&msg.User)
	if ok {
		msg.User.Status = ok
	}

	msg.User.Status = ok
	fmt.Println("write in channel")
	chanOut <- msg
}

func (u *UserService) LoginUser(chanOut chan *serviceModels.MessageOut) {
	msg := <-chanOut
	ok := u.UserDBService.LoginUser(&msg.User)
	if ok {
		u.UserDBService.GetUser(&msg.User)
		u.UserDBService.GetContactList(msg.User.Username)
		u.GroupDBService.GetGroupList(msg.User.Username)
	}
}

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
