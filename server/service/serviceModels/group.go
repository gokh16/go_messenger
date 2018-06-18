package serviceModels

type Group struct {
	GroupName string
	GroupType uint
	Members   []User
	Messages  []Message
}
