package interfaces

import "go_messenger/server/models"

//UserManager as contract between ORM level and Service Level
type UserManager interface {
	CreateUser(user *models.User) (bool, error)
	LoginUser(user *models.User) (bool, error)
	AddContact(user, contact *models.User, relationType uint) (bool, error)
	GetUsers(users *[]models.User)
	GetUser(user *models.User) (models.User, error)
	GetAccount(user *models.User) (models.User, error)
	GetContactList(user *models.User) ([]models.User, error)
	EditUser(user *models.User) models.User
	DeleteUser(user *models.User) (bool, error)
	DeleteContact(user, contact *models.User) (bool, error)
}
