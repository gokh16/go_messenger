package models

import (
	"github.com/jinzhu/gorm"
)

type MessageTemp struct {
	gorm.Model

	User               User
	Group              Group
	MessageContentType MessageContentType

	Content       string `json:"message_content"`
	MessageSender string `json:"username"`
	Email         string `json:"email"`

	MessageRecipientID   uint
	MessageContentTypeID uint
}
