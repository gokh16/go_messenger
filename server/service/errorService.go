package service

import (
	"go_messenger/server/models"
	"go_messenger/server/service/serviceModels"
	"log"
)

//ErrorService is a struct with an error
type ErrorService struct {
	err error
}

//Error - getter for error
func (e *ErrorService) Error() string {
	return e.err.Error()
}

//The SendError method sends description of an error to the client
func (e *ErrorService) SendError(err error, recipient models.User, chanOut chan<- *serviceModels.MessageOut) {
	e.err = err
	messageOut := serviceModels.MessageOut{User: recipient, Action: "Error"}
	messageOut.Err = e.Error()
	log.Println(messageOut.Err)
	chanOut <- &messageOut
}
