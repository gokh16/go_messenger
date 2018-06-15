package db_Interfaces

import "go_messenger/server/models"

type MessageDBI interface {
	AddMessage(content, username, groupName, contentType string) bool
	GetGroupMessages(groupName string) []models.Message
}
