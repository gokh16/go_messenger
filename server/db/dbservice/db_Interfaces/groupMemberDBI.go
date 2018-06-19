package db_Interfaces

import "go_messenger/server/models"

type GroupMemberDBI interface {
	AddGroupMember(userName, groupName, lastMessage string) bool
	GetGroupUserList(groupName string) []models.User
}
