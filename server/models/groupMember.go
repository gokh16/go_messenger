package models

import (
	"github.com/jinzhu/gorm"
)

//GroupMembers is a model to Database table
type GroupMember struct {
	gorm.Model

	User    User
	Group   Group
	Message Message

	UserID            uint
	GroupID           uint
	LastReadMessageID uint
}
