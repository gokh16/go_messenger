package dbservice

import "go_messenger/server/models"

//Message struct
type Message struct {
	models.Message
}

//AddMessage func
func (msg Message) AddMessage(content, username, groupName string, contentType uint) bool {
	sender := models.User{}
	recipient := models.Group{}
	conn.Where("username = ?", username).First(&sender)
	conn.Where("group_name = ?", groupName).First(&recipient)
	message := models.Message{Content: content, MessageSenderID: sender.ID, MessageRecipientID: recipient.ID, MessageContentTypeID: contentType}
	if conn.NewRecord(message) {
		conn.Create(&message)
		return true
	}
	return false
}
