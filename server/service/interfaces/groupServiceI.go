package interfaces

import (
	"go_messenger/server/userConnections"
	"go_messenger/server/service/serviceModels"
)

//groupInterface as contract between ORM level and Service Level
type GroupServiceI interface {
	CreateGroup(message *userConnections.MessageIn, chanOut chan *serviceModels.MessageOut)
	GetGroup(message *userConnections.MessageIn, chanOut chan *serviceModels.MessageOut)
	GetGroupList(message *userConnections.MessageIn, chanOut chan *serviceModels.MessageOut)
	EditGroup(message *userConnections.MessageIn, chanOut chan *serviceModels.MessageOut)
	AddGroupMember(message *userConnections.MessageIn, chanOut chan *serviceModels.MessageOut)
}
