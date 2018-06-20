package serviceModels

import "go_messenger/server/models"

//MessageOut responce struct
type MessageOut struct {
	User        models.User
	Members     []models.User
	ContactList []models.User
	GroupList   []Group
	Status      bool
	Action      string
	Err         error
}