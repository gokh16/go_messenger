<<<<<<< HEAD
package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Login    string
	Password string
	Username string `json:"username"`
	Email    string `json:"email"`
	Status   bool
	UserIcon string
}
=======
package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Login    string
	Password string
	Username string `json:"username"`
	Email    string `json:"email"`
	Status   bool
	UserIcon string
}
>>>>>>> 3661ec18fda6f6db02155e9be22dd834f0e1cd48
