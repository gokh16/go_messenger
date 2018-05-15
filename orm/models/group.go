package models

import (
	"github.com/jinzhu/gorm"
)

type Group struct {
	gorm.Model

	User User
	Group_Type Group_Type
	
	Group_name string
	Group_ownerID uint
	Group_typeID uint 
}
