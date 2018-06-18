package serviceModels

import "go_messenger/server/models"

type Group struct {
	GroupName string
	GroupType uint
	Members   []models.User    //??
	Messages  []models.Message //??
}
