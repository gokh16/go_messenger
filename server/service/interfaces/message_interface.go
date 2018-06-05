package interfaces

//messageInterface interface
type messageInterface interface {
	AddMessage(content, username, groupName string, contentType uint) bool
}

//MI is the type of messageInterface
type MI messageInterface
