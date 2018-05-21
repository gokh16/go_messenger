package models

import (
	"github.com/jinzhu/gorm"
)

type Message struct {
	gorm.Model

	User               User
	Group              Group
	MessageContentType MessageContentType

	Content              string `json:"message_content"`
	MessageSenderID      uint   `json:"message_sender_id"`
	MessageRecipientID   uint   `json:"message_recepient_id"`
	MessageContentTypeID uint   `json:"message_content_type_id"`
}
