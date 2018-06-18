package dbservice

import "go_messenger/server/models"

//MessageDBService struct
type MessageDBService struct {
	models.Message
}

//AddMessage func.
func (msg *MessageDBService) AddMessage(content, username, groupName, contentType string) bool {
	sender := models.User{}
	recipient := models.Group{}
	dbConn.Where("username = ?", username).First(&sender)
	dbConn.Where("group_name = ?", groupName).First(&recipient)
	message := models.Message{Content: content, MessageSenderID: sender.ID, MessageRecipientID: recipient.ID}
	if dbConn.NewRecord(message) {
		dbConn.Create(&message)
		return true
	}
	return false
}
func (msg *MessageDBService) GetGroupMessages(groupName string) []models.Message {
	var messageList = []models.Message{}
	return messageList
}