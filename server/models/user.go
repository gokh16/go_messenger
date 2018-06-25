package models

import (
	"github.com/jinzhu/gorm"
)

//User is a model to Database table
type User struct {
	gorm.Model

	Login    string
	Password string
	Username string `json:"username"`
	Email    string `json:"email"`
	Status   bool
	UserIcon string
}