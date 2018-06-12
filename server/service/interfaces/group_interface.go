package interfaces

import (
	"go_messenger/server/models"
)

//groupInterface as contract between ORM level and Service Level
type groupInterface interface {
	CreateGroup(groupName, groupOwner string, groupType uint) bool
	GetGroupList(userName string) []models.Group
	GetGroup(groupName string) models.Group
	//Delete(group *models.Group)
}

//GI is the type of groupInterface
type GI groupInterface
