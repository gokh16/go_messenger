package dbservice

import (
	"go_messenger/server/models"
)

//GroupMember type with build-in model of Group.
type GroupMember struct {
	models.GroupMember
}

//AddGroupMember method creates new record in DB GroupMember table with using the gorm framework. It returns bool value.
func (gm GroupMember) AddGroupMember(username, groupName, lastmessage string) bool {
	user := models.User{}
	group := models.Group{}
	message := models.Message{}
	conn.Where("username = ?", username).First(&user)
	conn.Where("group_name = ?", groupName).First(&group)
	conn.Where("content = ?", lastmessage).First(&message)
	member := models.GroupMember{UserID: user.ID, GroupID: group.ID, LastReadMessageID: message.ID}
	if conn.NewRecord(member) {
		conn.Create(&member)
		return true
	}
	return false
}

//GetGroupUserList gets all users of specific group and returns slice.
func (gm GroupMember) GetGroupUserList(groupName string) []models.User {
	group := models.Group{}
	users := []models.User{}
	conn.Where("group_name = ?", groupName).First(&group)
	conn.Joins("join group_members on users.id=group_members.user_id").Where("group_id =?", group.ID).Find(&users)
	return users
}
