package db_Interfaces

import "go_messenger/server/models"

type GroupDBI interface {
	CreateGroup(groupName, groupOwner string, groupType uint) bool
	GetGroupList(userName string) []models.Group
	GetGroup(groupName string) models.Group
	AddGroupMember(userName, groupName string, lastMessage uint) bool
	GetGroupUserList(groupName string) []models.User
}
