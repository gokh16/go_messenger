package dbservice

import (
	"errors"
	"go_messenger/server/models"
	"log"
)

//GroupDBService type with build-in model of Group.
type GroupDBService struct {
	models.Group
}

//CreateGroup method creates new record in DB Group table.
// It returns bool value.
func (g *GroupDBService) CreateGroup(group *models.Group) (bool, error) {
	dbConn.Where("login = ?", group.User.Login).First(&group.User)
	dbConn.Where("group_name = ?", group.GroupName).First(&group)
	group.GroupOwnerID = group.User.ID
	if dbConn.NewRecord(group) {
		dbConn.Create(&group)
		return true, dbConn.Error
	}
	return false, dbConn.Error
}

//AddGroupMember method creates new record in DB GroupMember table.
// It returns bool value.
func (g *GroupDBService) AddGroupMember(user *models.User, group *models.Group, message *models.Message) (bool, error) {
	dbConn.Where("login = ?", user.Login).Take(&user)
	dbConn.Where("group_name = ?", group.GroupName).First(&group)
	dbConn.Where("content = ?", message.Content).First(&message)
	member := models.GroupMember{UserID: user.ID, GroupID: group.ID, LastReadMessageID: message.ID}
	if dbConn.NewRecord(member) {
		dbConn.Create(&member)
		return true, dbConn.Error
	}
	return false, dbConn.Error
}

//GetGroupList method gets all groups of special user from DB.
// It returns slice []models.Group.
func (g *GroupDBService) GetGroupList(user *models.User) ([]models.Group, error) {
	var groupList []models.Group
	dbConn.Where("login = ?", user.Login).First(&user)
	dbConn.Joins("join group_members on groups.id=group_members.group_id").Where("user_id = ?", user.ID).Find(&groupList)
	return groupList, dbConn.Error
}

//GetGroup method gets group of special user from DB.
//It returns object of models.Group.
func (g *GroupDBService) GetGroup(group *models.Group) (models.Group, error) {
	record := dbConn.Where("group_name = ?", group.GroupName).Take(&group)
	switch {
	case record.RecordNotFound():
		err := errors.New("Group not found")
		return *group, err
	case record.Error != nil:
		return *group, record.Error
	default:
		return *group, nil
	}
}

//GetMemberList method gets all members of special group from DB.
// It returns slice []models.User.
func (g *GroupDBService) GetMemberList(group *models.Group) ([]models.User, error) {
	memberList := []models.User{}
	dbConn.Where("group_name = ?", group.GroupName).First(&group)
	dbConn.Joins("join group_members on users.id=group_members.user_id").Where("group_id =?", group.ID).Find(&memberList)
	return memberList, dbConn.Error
}

//EditGroup method updates the relevant entry in the DB
//It returns bool value
func (g *GroupDBService) EditGroup(group *models.Group) bool {
	var groupInstance models.Group
	dbConn.Where("id = ?", group.ID).Take(&groupInstance)
	if group.GroupName != "" {
		groupInstance.GroupName = group.GroupName
		dbConn.Save(&groupInstance)
		log.Printf("EDIT GROUP: ID %d, Group %s was updated on %s", groupInstance.ID, groupInstance.GroupName, group.GroupName)
		return true
	}
	return false
}
