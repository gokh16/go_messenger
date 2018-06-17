package serviceModels

type User struct {
	Users    []User
	Login    string
	Username string
	Email    string
	Status   bool
	UserIcon string
}
