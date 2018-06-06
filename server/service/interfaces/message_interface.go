package interfaces

//messageInterface interface
type messageInterface interface {
<<<<<<< HEAD
	AddMessage(content, username, groupName, contentType string) bool
=======
	AddMessage(content, username, groupName string, contentType string) bool
>>>>>>> group-chat
}

//MI is the type of messageInterface
type MI messageInterface
