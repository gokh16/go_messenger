package interfaces

//userInterface as contract between ORM level and Service Level
type userInterface interface {
	CreateUser(login, password, username, email string, status bool, usericon string) bool
	//Delete(user *models.User)
}

//UI is the type of userInterface
type UI userInterface
