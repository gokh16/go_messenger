package interfaces

import (
	"go_messenger/server/userConnections"
)

//groupInterface as contract between ORM level and Service Level
type GroupServiceI interface {
	CreateGroup(message *userConnections.Message, chanOut chan *userConnections.Message)
	GetGroup(message *userConnections.Message, chanOut chan *userConnections.Message)
	GetGroupList(message *userConnections.Message, chanOut chan *userConnections.Message)
	EditGroup(message *userConnections.Message, chanOut chan *userConnections.Message)
	AddGroupMember(message *userConnections.Message, chanOut chan *userConnections.Message)
}
