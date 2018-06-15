package dbservice

import (
	"go_messenger/server/models"
)

//User type with build-in model of User.
type UserDBService struct {
	models.User
}

//CreateUser method creates record in DB with using the gorm framework. It returns bool value.
func (u *UserDBService) CreateUser(user *models.User) bool {
	dbConn.Where("username = ?", user.Username).First(&user)
	if dbConn.NewRecord(user) {
		dbConn.Create(&user)
		return true
	}
	return false
}

//CreateUser method creates record in DB with using the gorm framework. It returns bool value.
func (u *UserDBService) GetUsers(users *[]models.User) {
	dbConn.Find(&users)
}

//LoginUser method get record from DB with using the gorm framework. It returns bool value.
func (u *UserDBService) LoginUser(user *models.User) bool {
	dbConn.Where("login = ?", user.Login).Where("password = ?", user.Password).First(&user)
	if dbConn.NewRecord(user) {
		return false
	}
	return true
}

//GetUser method get record from DB with using the gorm framework. It returns User object.
func (u *UserDBService) GetUser(user *models.User) models.User {
	dbConn.Where("username = ?", user.Username).First(&user)
	return *user
}

func (u *UserDBService) AddContact(userName, contactName string, relationType uint) bool {
	user := models.User{}
	contact := models.User{}
	dbConn.Where("username = ?", userName).First(&user)
	dbConn.Where("username = ?", contactName).First(&contact)
	relation := models.UserRelation{RelatingUser: user.ID, RelatedUser: contact.ID, RelationTypeID: relationType}
	if dbConn.NewRecord(relation) {
		dbConn.Create(&relation)
		return true
	}
	return false

}

func (u *UserDBService) GetContactList(userName string) []models.User {
	user := models.User{}
	contactList := []models.User{}
	temp := []models.UserRelation{}
	dbConn.Where("username = ?", userName).First(&user)
	dbConn.Where("relating_user=?", user.ID).Find(&temp)
	for i, _ := range temp {
		contact := models.User{}
		dbConn.Where("id=?", temp[i].RelatedUser).First(&contact)
		contactList = append(contactList, contact)

	}
	return contactList
}
