package db_Interfaces

import "go_messenger/server/models"

type GroupDBI interface {
	CreateGroup(groupName, groupOwner string, groupType uint) bool
	GetGroupList(userName string) []models.Group
	GetGroup(groupName string) models.Group
	AddGroupMember(username, groupName, lastmessage string) bool
	GetGroupUserList(groupName string) []models.User
}
