package interfaces

import "go_messenger/server/models"

//userInterface as contract between ORM level and Service Level
type userInterface interface {
	CreateUser(user *models.User) bool
	LoginUser(user *models.User) bool
	GetUser(user *models.User) models.User
	GetContactList() []models.User
	GetUsers(users *[]models.User)
	//Delete(user *models.User)
}

//UI is the type of userInterface
type UI userInterface
