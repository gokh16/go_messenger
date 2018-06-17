package dbservice

import (
	"go_messenger/server/models"
)

//Group type with build-in model of Group.
type Group struct {
	models.Group
}

//CreateGroup method creates new record in DB Group table with using the gorm framework. It returns bool value.
func (g Group) CreateGroup(groupName, groupOwner string, groupType uint) bool {
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

//GetGroupList is getting users from DB
func (g Group) GetGroupList(userName string) []models.Group {
	user := models.User{}
	groups := []models.Group{}
	dbConn.Where("username = ?", userName).First(&user)
	dbConn.Joins("join group_members on groups.id=group_members.group_id").Where("user_id = ?", user.ID).Find(&groups)
	return groups
}

func (g Group) GetGroup(groupName string) models.Group {
	group := models.Group{}
	dbConn.Where("groupname = ?", groupName).First(&group)
	return group
}

//AddGroupMember method creates new record in DB GroupMember table with using the gorm framework. It returns bool value.
func (gm Group) AddGroupMember(username, groupName, lastmessage string) bool {
	user := models.User{}
	group := models.Group{}
	message := models.Message{}
	dbConn.Where("username = ?", username).First(&user)
	dbConn.Where("group_name = ?", groupName).First(&group)
	dbConn.Where("content = ?", lastmessage).First(&message)
	member := models.GroupMember{UserID: user.ID, GroupID: group.ID, LastReadMessageID: message.ID}
	if dbConn.NewRecord(member) {
		dbConn.Create(&member)
		return true
	}
	return false
}

//GetGroupUserList gets all users of specific group and returns slice.
func (gm Group) GetGroupUserList(groupName string) []models.User {
	group := models.Group{}
	users := []models.User{}
	dbConn.Where("group_name = ?", groupName).First(&group)
	dbConn.Joins("join group_members on users.id=group_members.user_id").Where("group_id =?", group.ID).Find(&users)
	return users
}
