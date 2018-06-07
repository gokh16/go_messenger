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

//LoginUser method get record from DB with using the gorm framework. It returns bool value.
func (u User) LoginUser(user *models.User) bool {
	dbConn.Where("login = ?", user.Login).Where("password = ?", user.Password).First(&user)
	if dbConn.NewRecord(user) {
		return false
	}
	return true
}

//GetUser method get record from DB with using the gorm framework. It returns User object.
func (u User) GetUser(user *models.User) models.User {
	dbConn.Where("login = ?", user.Login).First(&user)
	return *user
}

func (u User) GetContactList() []models.User {
	userList := []models.User{}
	return userList
}
