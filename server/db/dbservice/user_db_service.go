package dbservice

import (
	"go_messenger/server/models"
)

//User type with build-in model of User.
type User struct {
	models.User
}

//CreateUser method creates record in DB with using the gorm framework. It returns bool value.
func (u User) CreateUser(user *models.User) bool {
	dbConn.Where("username = ?", user.Username).First(&user)
	if dbConn.NewRecord(user) {
		dbConn.Create(&user)
		return true
	}
	return false
}
