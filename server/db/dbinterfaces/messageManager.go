package dbinterfaces

import "go_messenger/server/models"

//MessageManager interface
type MessageManager interface {
	AddMessage(content, username, groupName, contentType string) bool
	GetGroupMessages(groupName string) []models.Message
}
