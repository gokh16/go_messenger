package dbservice

import (
	"go_messenger/server/models"
)

//GroupMember type with build-in model of Group.
type GroupMemberDBService struct {
	models.GroupMember
}

//AddGroupMember method creates new record in DB GroupMember table with using the gorm framework. It returns bool value.
//TODO M
func (gm *GroupMemberDBService) AddGroupMember(username, groupName, lastmessage string) bool {
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
func (gm *GroupMemberDBService) GetGroupUserList(groupName string) []models.User {
	group := models.Group{}
	users := []models.User{}
	dbConn.Where("group_name = ?", groupName).First(&group)
	dbConn.Joins("join group_members on users.id=group_members.user_id").Where("group_id =?", group.ID).Find(&users)
	return users
}
