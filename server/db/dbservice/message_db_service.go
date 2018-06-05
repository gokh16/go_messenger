package dbservice

import "go_messenger/server/models"

//Message struct
type Message struct {
	models.Message
}

//AddMessage func
func (msg Message) AddMessage(content, username, groupName string, contentType string) bool {
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
