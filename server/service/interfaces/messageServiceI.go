package interfaces

import (
	"go_messenger/server/userConnections"
	"go_messenger/server/service/serviceModels"
)

//messageInterface interface
type MessageServiceI interface {
	SendMessageTo(message *userConnections.MessageIn, chanOut chan *serviceModels.MessageOut)
}

