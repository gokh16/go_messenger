package structure

import (
	"github.com/jinzhu/gorm"
)

//User is a model to Database table
type User struct {
	gorm.Model

	Login    string
	Password string
	Username string
	Email    string
	Status   bool
	UserIcon string
}