package db_Interfaces

import "go_messenger/server/models"

type UserDBI interface {
	CreateUser(user *models.User) bool
	GetUsers(users *[]models.User)
	LoginUser(user *models.User) bool
	GetUser(user *models.User) models.User
	AddContact(userName, contactName string, relationType uint) bool
	GetContactList(userName string) []models.User
}
