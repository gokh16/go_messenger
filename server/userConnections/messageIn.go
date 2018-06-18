package userConnections

import "go_messenger/server/models"

//Message is a main structure which contains fields for interactions client and server
type MessageIn struct {
	User      models.User
	Group     models.Group
	GroupType models.GroupType
	Message   models.Message
	Action    string
}
