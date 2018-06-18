package serviceModels

//MessageOut ...
type MessageOut struct {
	User        User
	ContactList []User
	GroupList   []Group
	Status      bool
	Action      string
	Err         error
}
