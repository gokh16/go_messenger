package userConnections

import "go_messenger/server/models"

type Message struct {
	User        models.User
	Group       models.Group
	GroupMember models.GroupMember
	Member      Member
	Message     models.Message
	Status      bool
	Action      string
}

type Member struct {
	GroupMembers []models.User
}
