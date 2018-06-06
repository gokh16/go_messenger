package interfaces

import "go_messenger/server/models"

//userInterface as contract between ORM level and Service Level
type userInterface interface {
	CreateUser(user *models.User) bool
	//Delete(user *models.User)
}

//UI is the type of userInterface
type UI userInterface
