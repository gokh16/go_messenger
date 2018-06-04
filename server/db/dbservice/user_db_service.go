package dbservice

import (
	"go_messenger/server/models"
)

//User type with build-in model of User.
type User struct {
	models.User
}

//CreateUser method creates record in DB with using the gorm framework. It returns bool value.
func (u User) CreateUser(login, password, username, email string, status bool, usericon string) bool {
	user := models.User{Login: login, Password: password, Username: username, Email: email, Status: status, UserIcon: usericon}
	if dbConn.NewRecord(user) {
		dbConn.Create(&user)
		return true
	}
	return false
}
