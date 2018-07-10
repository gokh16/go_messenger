package serviceModels

import "go_messenger/server/models"

//MessageOut response struct
type MessageOut struct {
	User        models.User
	Recipients  []models.User
	Message     models.Message
	ContactList []models.User
	GroupList   []Group
	Status      bool
	Action      string
	Err         string
}