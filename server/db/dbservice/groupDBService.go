package dbservice

import (
	"go_messenger/server/models"
)

//GroupDBService type with build-in model of GroupDBService.
type GroupDBService struct {
	models.Group
}

//CreateGroup method creates new record in DB GroupDBService table with using the gorm framework. It returns bool value.
func (g *GroupDBService) CreateGroup(groupName string, groupOwner, groupType uint) bool {
	owner := models.User{}
	dbConn.Where("username = ?", groupOwner).First(&owner)
	group := models.Group{GroupName: groupName, GroupOwnerID: owner.ID, GroupTypeID: groupType}
	dbConn.Where("groupname = ?", groupName).First(&group)
	if dbConn.NewRecord(group) {
		dbConn.Create(&group)
		return true
	}
	return false
}

func (g *GroupDBService) GetGroupList(userName string) []models.Group {
	user := models.User{}
	groups := []models.Group{}
	dbConn.Where("username = ?", userName).First(&user)
	dbConn.Joins("join group_members on groups.id=group_members.group_id").Where("user_id = ?", user.ID).Find(&groups)
	return groups
}

func (g *GroupDBService) GetGroup(groupName string) models.Group {
	group := models.Group{}
	dbConn.Where("groupname = ?", groupName).First(&group)
	return group
}

//AddGroupMember method creates new record in DB GroupMembers table with using the gorm framework. It returns bool value.
func (g *GroupDBService) AddGroupMember(userName, groupName string, lastMessage uint) bool {
	user := models.User{}
	group := models.Group{}
	message := models.Message{}
	dbConn.Where("userName = ?", userName).First(&user)
	dbConn.Where("group_name = ?", groupName).First(&group)
	dbConn.Where("content = ?", lastMessage).First(&message)
	member := models.GroupMember{UserID: user.ID, GroupID: group.ID, LastReadMessageID: message.ID}
	if dbConn.NewRecord(member) {
		dbConn.Create(&member)
		return true
	}
	return false
}

//GetGroupUserList gets all users of specific group and returns slice.
func (g *GroupDBService) GetGroupUserList(groupName string) []models.User {
	group := models.Group{}
	users := []models.User{}
	dbConn.Where("group_name = ?", groupName).First(&group)
	dbConn.Joins("join group_members on users.id=group_members.user_id").Where("group_id =?", group.ID).Find(&users)
	return users
}