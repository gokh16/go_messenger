package structure

import (
	"github.com/jinzhu/gorm"
)

//Message is a model to Database table
type Message struct {
	gorm.Model

	User  User
	Group Group

	Content            string
	MessageSenderID    uint
	MessageRecipientID uint
	MessageContentType string
}