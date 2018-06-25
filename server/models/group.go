package models

import (
	"github.com/jinzhu/gorm"
)

//Group is a model to Database table.
type Group struct {
	gorm.Model

	User      User
	GroupType GroupType

	GroupName    string
	GroupOwnerID uint
	GroupTypeID  uint
}
