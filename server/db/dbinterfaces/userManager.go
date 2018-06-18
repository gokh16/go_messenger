package dbinterfaces

import "go_messenger/server/models"

//userInterface as contract between ORM level and Service Level
type UserManager interface {
	CreateUser(user *models.User) bool
	LoginUser(user *models.User) bool
	GetUser(user *models.User) models.User
	GetContactList(user *models.User) []models.User
	GetUsers(users *[]models.User)
	//Delete(user *models.User)
}
