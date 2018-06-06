package interfaces

//messageInterface interface
type messageInterface interface {
	AddMessage(content, username, groupName, contentType string) bool
}

//MI is the type of messageInterface
type MI messageInterface
