package dbservice

import (
	"go_messenger/server/models"
	"log"
)

//Group type with build-in model of Group.
type GroupDBService struct {
	models.Group
}

//CreateGroup method creates new record in DB Group table.
// It returns bool value.
func (g GroupDBService) CreateGroup(group *models.Group) bool {
	log.Println(group.User.Username, group.GroupName)
	dbConn.Where("username = ?", group.User.Username).First(&group.User)
	dbConn.Where("group_name = ?", group.GroupName).First(&group)
	group.GroupOwnerID = group.User.ID
	if dbConn.NewRecord(group) {
		dbConn.Create(&group)
		return true
	}
	return false
}

//AddGroupMember method creates new record in DB GroupMember table.
// It returns bool value.
func (g GroupDBService) AddGroupMember(user *models.User, group *models.Group, message *models.Message) bool {
	dbConn.Where("username = ?", user.Username).First(&user)
	dbConn.Where("group_name = ?", group.GroupName).First(&group)
	dbConn.Where("content = ?", message.Content).First(&message)
	member := models.GroupMember{UserID: user.ID, GroupID: group.ID, LastReadMessageID: message.ID}
	if dbConn.NewRecord(member) {
		dbConn.Create(&member)
		return true
	}
	return false
}

//GetGroupList method gets all groups of special user from DB.
// It returns slice []models.Group.
func (g GroupDBService) GetGroupList(user *models.User) []models.Group {
	var groupList []models.Group
	dbConn.Where("login = ?", user.Login).First(&user)
	dbConn.Joins("join group_members on groups.id=group_members.group_id").Where("user_id = ?", user.ID).Find(&groupList)
	return groupList
}

//GetGroup method gets group of special user from DB.
// It returns object of models.Group.
func (g GroupDBService) GetGroup(group *models.Group) models.Group {
	//dbConn.Where("username = ?", userName).First(&user)  ????
	dbConn.Where("group_name = ?", group.GroupName).Where("user_id = ?", group.User.ID).Find(&group)
	return *group
}

//GetMemberList method gets all members of special group from DB.
// It returns slice []models.User.
func (g GroupDBService) GetMemberList(group *models.Group) []models.User {
	memberList := []models.User{}
	dbConn.Where("group_name = ?", group.GroupName).First(&group)
	dbConn.Joins("join group_members on users.id=group_members.user_id").Where("group_id =?", group.ID).Find(&memberList)
	return memberList
}