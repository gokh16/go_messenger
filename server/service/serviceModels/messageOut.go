package serviceModels

import "go_messenger/server/models"

type MessageOut struct {
	User   models.User     //?? maiby not models
	Users  []models.User
	Groups []Group
	Status bool
	Action string
	Err    error
}
