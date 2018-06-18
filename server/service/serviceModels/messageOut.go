package serviceModels

type Group struct {
	GroupName string
	GroupType uint
	Members   []User
	Messages  []Message
}

//User for MessageOut
type User struct {
	Login    string
	Username string
	Email    string
	UserIcon string
}

//Message for MessageOut
type Message struct {
	Content string
}

//MessageOut ...
type MessageOut struct {
	User        User
	ContactList []User
	GroupList   []Group
	Status      bool
	Action      string
	Err         error
}
