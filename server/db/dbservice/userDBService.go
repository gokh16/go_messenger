package dbservice

import (
	"errors"
	"go_messenger/server/models"
)

//UserDBService type with build-in model of User.
type UserDBService struct {
	models *models.User
}

//CreateUser method creates User in DB.
//It returns bool value.
func (u *UserDBService) CreateUser(user *models.User) (bool, error) {
	dbConn.Where("login = ?", user.Login).First(&user)
	if dbConn.NewRecord(user) {
		dbConn.Create(&user)
		return true, dbConn.Error
	}
	return false, dbConn.Error
}

//LoginUser - user's auth.
func (u *UserDBService) LoginUser(user *models.User) (bool, error) {
	record := dbConn.Where("password = ?", user.Password).Where("login = ?", user.Login).Take(&user)
	switch {
	case dbConn.Error != nil:
		return false, dbConn.Error
	case record.RecordNotFound():
		err := errors.New("Login or Password does not correct")
		return false, err
	default:
		return true, dbConn.Error

	}
}

func (u *UserDBService) GetAccount(user *models.User) (models.User, error) {
	dbConn.Where("login = ?", user.Login).Take(&user)
	if dbConn.Error != nil {
		return *user, dbConn.Error
	}
	return *user, dbConn.Error
}

//AddContact add spesial user to contact list of special User
func (u *UserDBService) AddContact(user, contact *models.User, relationType uint) (bool, error) {
	dbConn.Where("login = ?", user.Login).First(&user)
	dbConn.Where("login = ?", contact.Login).First(&contact)
	relation := models.UserRelation{RelatingUser: user.ID, RelatedUser: contact.ID, RelationTypeID: relationType}
	if dbConn.NewRecord(relation) {
		dbConn.Create(&relation)
		return true, dbConn.Error
	}
	return false, dbConn.Error

}

//GetUsers method gets all users from DB.
func (u *UserDBService) GetUsers(users *[]models.User) {
	dbConn.Find(&users)
}

//GetUser method get special user from DB.
//It returns User object.
func (u *UserDBService) GetUser(user *models.User) (models.User, error) {
	record := dbConn.Where("login = ?", user.Login).Take(&user)
	switch {
	case dbConn.Error != nil:
		return *user, dbConn.Error
	case record.RecordNotFound():
		err := errors.New("User does not exist")
		return *user, err
	default:
		return *user, nil
	}
}

//GetContactList gets contact list of special user from DB.
//It returns slice []models.User.
func (u *UserDBService) GetContactList(user *models.User) ([]models.User, error) {
	contactList := []models.User{}
	temp := []models.UserRelation{}
	dbConn.Where("login = ?", user.Login).First(&user)
	dbConn.Where("relating_user=?", user.ID).Find(&temp)
	for i := range temp {
		contact := models.User{}
		dbConn.Where("id=?", temp[i].RelatedUser).First(&contact)
		contactList = append(contactList, contact)
	}
	return contactList, dbConn.Error
}

//DeleteUser delete account from DB.
//It returns bull and error.
func (u *UserDBService) DeleteUser(user *models.User) (bool, error) {
	record := dbConn.Where("login = ?", user.Login).Take(&user)
	if record.RecordNotFound() {
		err := errors.New("User already deleted")
		return false, err
	}
	record.Delete(&user)
	if record.Error != nil {
		return false, record.Error
	}
	return true, nil
}
