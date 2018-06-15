package interfaces

import (
	"go_messenger/server/userConnections"
)

//messageInterface interface
type MessageServiceI interface {
	SendMessageTo(message *userConnections.Message, chanOut chan *userConnections.Message)
}

