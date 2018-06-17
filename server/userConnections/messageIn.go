package userConnections

import "go_messenger/server/models"

type MessageIn struct {
	User      models.User
	Group     models.Group
	GroupType models.GroupType
	Message   models.Message
	Action    string
}
