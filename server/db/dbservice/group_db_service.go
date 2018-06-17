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
