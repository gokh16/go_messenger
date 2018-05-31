<<<<<<< HEAD
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
=======
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
>>>>>>> 3661ec18fda6f6db02155e9be22dd834f0e1cd48
