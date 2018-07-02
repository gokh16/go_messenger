package util

import (
	"go_messenger/desktop/structure"
)

func NewUser(login string, password string, username string, email string, status bool, userIcon string) *structure.User {
	return &structure.User{
		Login:    login,
		Password: password,
		Username: username,
		Email:    email,
		Status:   status,
		UserIcon: userIcon,
	}
}

func NewGroup(user *structure.User, groupType string, groupName string, groupOwnerID uint, groupTypeID uint) *structure.Group {
	return &structure.Group{
		User:         *user,
		GroupName:    groupName,
		GroupTypeID:  groupTypeID,
		GroupOwnerID: groupOwnerID,
		GroupType:    structure.GroupType{Type: groupType},
	}
}

func NewMessage(user *structure.User, group *structure.Group, content string, messageSenderID uint, messageRecipientID uint, messageContentType string) *structure.Message {
	return &structure.Message{
		User:               *user,
		Group:              *group,
		Content:            content,
		MessageSenderID:    messageSenderID,
		MessageRecipientID: messageRecipientID,
		MessageContentType: messageContentType,
	}
}

func NewMessageOut(user *structure.User, contact *structure.User, group *structure.Group, message *structure.Message, members []structure.User, relationType uint, messageLimit uint, action string) *MessageOut {
	return &MessageOut{
		User:         *user,
		Contact:      *contact,
		Group:        *group,
		Message:      *message,
		Members:      members, //todo pointer??
		RelationType: relationType,
		MessageLimit: messageLimit,
		Action:       action,
	}
}
