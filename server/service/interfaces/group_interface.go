package interfaces

import (
	"go_messenger/server/models"
)

//groupInterface as contract between DBservice level and Service Level
type groupInterface interface {
	CreateGroup(groupName, groupOwner string, groupType uint) bool
	GetGroupList(userName string) []models.Group
	GetGroup(userName, groupName string) models.Group
	AddGroupMember(username, groupName, lastmessage string) bool
	GetGroupUserList(groupName string) []models.User
	//Delete(group *models.Group)
}

//GI is the type of groupInterface
type GI groupInterface
