package userConnections

import "go_messenger/server/models"

type Message struct {
	UserName     string
	RelatingUser string
	RelatedUser  string
	RelationType uint
	GroupName    string
	GroupType    uint
	GroupOwner   string
	GroupMember  []string
	ContentType  string
	Content      string
	LastMessage  string
	Login        string
	Password     string
	Email        string
	Status       bool
	UserIcon     string
	Action       string
}
type MessageIn struct {
	User      models.User
	Group     models.Group
	GroupType models.GroupType
	Message   models.Message
	Action    string
}
