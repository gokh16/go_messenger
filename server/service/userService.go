package service

import (
	"go_messenger/server/db/dbservice"
	"go_messenger/server/service/serviceModels"
	"go_messenger/server/userConnections"
)

type UserService struct {
	UserDBService dbservice.UserDBService
	GroupDBService dbservice.GroupDBService
}

func (u *UserService) GetUsers(msg *userConnections.MessageIn, chanOut chan *serviceModels.MessageOut) {
	//msgOut := serviceModels.MessageOut{}
	//u.UserDBService.GetUsers(&msgOut.Members)
	//msgOut.Action = msg.Action
	//chanOut <- &msgOut

}

func (u *UserService) CreateUser(msg *userConnections.MessageIn, chanOut chan *serviceModels.MessageOut) {
	//fmt.Println("Service Ok")
	//var ok bool
	//msgOut := serviceModels.MessageOut{}
	//if ok := u.UserDBService.CreateUser(&msg.User); !ok {
	//	msgOut.Status = ok
	//}
	//msgOut.Status = ok
	//msgOut.Action = msg.Action
	//fmt.Println("write in channel")
	//chanOut <- &msgOut
}

func (u *UserService) LoginUser(message *userConnections.MessageIn, chanOut chan *serviceModels.MessageOut) {
	//msg := <-chanOut
	//ok := u.UserDBService.LoginUser(&msg.User)
	//if ok {
	//	u.UserDBService.GetUser(&msg.User)
	//	u.UserDBService.GetContactList(msg.User.Username)
	//	u.GroupDBService.GetGroupList(msg.User.Username)
	//}
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
