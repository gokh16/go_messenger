package interfaces

import (
	"go_messenger/server/userConnections"
)

//userInterface as contract between ORM level and Service Level
type UserServiceI interface {
	CreateUser(message *userConnections.Message, chanOut chan *userConnections.Message)
	LoginUser(chanOut chan *userConnections.Message)
	GetUsers(message *userConnections.Message, chanOut chan *userConnections.Message)
}

