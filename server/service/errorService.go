package service

import (
	"errors"
	"go_messenger/server/service/serviceModels"
	"go_messenger/server/userConnections"
	"log"
)

type ErrorService struct {
	err error
}

func (e *ErrorService) Error() string {
	return "Unknown Action Error"
}

func (e *ErrorService) UnknownActionError(messageIn *userConnections.MessageIn, chanOut chan<- *serviceModels.MessageOut) {
	e.err = errors.New(e.Error())
	messageOut := serviceModels.MessageOut{User: messageIn.User, Action: "Error"}
	messageOut.Err = e.err.Error()

	log.Println(messageOut.Err)
	chanOut <- &messageOut
}
