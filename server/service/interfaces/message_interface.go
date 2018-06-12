package interfaces

import "go_messenger/server/models"

//messageInterface interface
type messageInterface interface {
	AddMessage(content, username, groupName, contentType string) bool
	GetGroupMessages(groupName string) []models.Message
}

//MI is the type of messageInterface
type MI messageInterface
