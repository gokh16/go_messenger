package dbinterfaces

import (
	"go_messenger/server/models"
)

//GroupManager as contract between DBservice level and Service Level
type GroupManager interface {
	CreateGroup(groupName, groupOwner string, groupType uint) bool
	GetGroupList(user *models.User) []models.Group
	GetGroup(userName, groupName string) models.Group
	AddGroupMember(username, groupName, lastmessage string) bool
	GetGroupUserList(groupName string) []models.User
	//Delete(group *models.Group)
}
