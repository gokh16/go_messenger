package models

import (
	"github.com/jinzhu/gorm"
)

type Group_Member struct {
	gorm.Model

	User User
	Group Group
	Message Message
	
	UserID uint 
	GroupID uint 
	Last_read_messageID uint 
}