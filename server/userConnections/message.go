package userConnections

import "go_messenger/server/models"

type Message struct {
	User        models.User
	Group       models.Group
	GroupMember GroupMember
	Message     models.Message
	Action      string
}

type GroupMember struct {
	GroupMember []models.User
}
