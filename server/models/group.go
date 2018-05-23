package models

import (
	"github.com/jinzhu/gorm"
)

type Group struct {
	gorm.Model

	User      User
	GroupType GroupType

	GroupName    string `json:"group_name"`
	GroupOwnerID uint   `json:"group_owner_id"`
	GroupTypeID  uint   `json:"group_type_id"`
}
