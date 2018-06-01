package models

import (
	"github.com/jinzhu/gorm"
)

type GroupMember struct {
	gorm.Model

	User    User
	Group   Group
	Message Message

	UserID            uint `json:"user_id"`
	GroupID           uint `json:"group_id"`
	LastReadMessageID uint `json:"last_read_message_id"`
}
