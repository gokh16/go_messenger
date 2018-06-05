package interfaces

import "go_messenger/server/models"

//groupMemberInterface as contract between ORM level and Service Level
type groupMemberInterface interface {
	AddGroupMember(username, groupName, lastmessage string) bool
	GetGroupUserList(groupName string) []models.User
	//DeleteGroupMember(groupMember *models.GroupMember)
}

//GMI is the type of groupMemberInterfase
type GMI groupMemberInterface
