package interfaces

import (
	"go_messenger/server/models"
)

//GroupManager as contract between DBservice level and Service Level
type GroupManager interface {
	CreateGroup(group *models.Group) bool
	AddGroupMember(user *models.User, group *models.Group, message *models.Message) bool
	GetGroupList(user *models.User) []models.Group
	GetGroup(group *models.Group) models.Group
	GetMemberList(group *models.Group) []models.User
}
