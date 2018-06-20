package serviceModels

import "go_messenger/server/models"

type Group struct {
	GroupName string
	GroupType models.GroupType
	Members   []models.User
	Messages  []models.Message
}