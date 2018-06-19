package dbservice

import "go_messenger/server/models"

//Message struct
type Message struct {
	models.Message
}

//AddMessage func.
func (msg Message) AddMessage(message *models.Message) bool {
	dbConn.Where("username = ?", message.User.Username).First(&message.User)
	dbConn.Where("group_name = ?", message.Group.GroupName).First(&message.Group)
	if dbConn.NewRecord(message) {
		dbConn.Create(&message)
		return true
	}
	return false
}
func (msg Message) GetGroupMessages(group *models.Group) []models.Message {
	var messageList = []models.Message{}

	return messageList
}
