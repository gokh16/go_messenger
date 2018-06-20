package userConnections

import "go_messenger/server/models"

//Message is a main structure which contains fields for interactions client and server
type MessageIn struct {
	User         models.User
	Contact      models.User
	Group        models.Group
	Message      models.Message
	Members      []models.User
	RelationType uint
	//show current limit of past messages
	MessageLimit uint
	Action       string
}
