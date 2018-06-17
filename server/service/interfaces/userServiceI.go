package interfaces

import (
	"go_messenger/server/userConnections"
	"go_messenger/server/service/serviceModels"
)

//userInterface as contract between ORM level and Service Level
type UserServiceI interface {
	CreateUser(message *userConnections.MessageIn, chanOut chan *serviceModels.MessageOut)
	LoginUser(chanOut chan *serviceModels.MessageOut)
	GetUsers(message *userConnections.MessageIn, chanOut chan *serviceModels.MessageOut)
}

