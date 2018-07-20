package serviceModels

import "go_messenger/server/models"

type Group struct {
	ID        uint
	GroupName string
	GroupType models.GroupType
	Members   []models.User
	Messages  []models.Message
}

func NewGroup(group models.Group, members []models.User, messages []models.Message) *Group {
	ID := group.ID
	GroupName := group.GroupName
	GroupType := group.GroupType
	Members := members
	Messages := messages
	return &Group{
		ID,
		GroupName,
		GroupType,
		Members,
		Messages,
	}

}
