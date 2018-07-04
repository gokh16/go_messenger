
package interfaces

import "go_messenger/server/models"

//MessageManager interface
type MessageManager interface {
	AddMessage(message *models.Message) bool
	GetGroupMessages(group *models.Group, count uint) []models.Message
}