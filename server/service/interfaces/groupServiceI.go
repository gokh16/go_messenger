package interfaces

import (
	"go_messenger/server/models"
)

//groupInterface as contract between ORM level and Service Level
type GroupI interface {
	CreateGroup(groupName, groupOwner string, groupType uint) bool
	GetGroupList(userName string) []models.Group
	GetGroup(groupName string) models.Group
	AddGroupMember(username, groupName, lastmessage string) bool
	GetGroupUserList(groupName string) []models.User
	//Delete(group *models.Group)
}
