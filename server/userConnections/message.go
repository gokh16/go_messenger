package userConnections

//Message is a main structure which contains fields for interactions client and server
type Message struct {
	UserName     string
	RelatingUser string
	RelatedUser  string
	RelationType uint
	GroupName    string
	GroupType    uint
	GroupOwner   string
	GroupMember  []string
	ContentType  string
	Content      string
	LastMessage  string
	Login        string
	Password     string
	Email        string
	Status       bool
	UserIcon     string
	Action       string
}
