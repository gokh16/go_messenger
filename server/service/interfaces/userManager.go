package interfaces

import "go_messenger/server/models"

//UserManager as contract between ORM level and Service Level
type UserManager interface {
	CreateUser(user *models.User) (bool, error)
	LoginUser(user *models.User) bool
	AddContact(user, contact *models.User, relationType uint) bool
	GetUsers(users *[]models.User)
	GetUser(user *models.User) models.User
	GetContactList(user *models.User) []models.User
}
