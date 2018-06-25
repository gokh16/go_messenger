package serviceModels

import "go_messenger/server/models"

//MessageOut responce struct
type MessageOut struct {
	User        models.User
	Members     []models.User
	Message     models.Message
	ContactList []models.User
	GroupList   []Group
	Status      bool
	Action      string
	Err         error
}