package models

import (
	"github.com/jinzhu/gorm"
)

type Message struct {
	gorm.Model

	User User
	Group Group
	Message_Content_Type Message_Content_Type
	
	Content string
	Message_senderID uint 
	Message_recipientID uint
	Message_content_typeID uint
}