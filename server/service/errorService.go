package service

import (
	"go_messenger/server/userConnections"
	"go_messenger/server/service/serviceModels"
	"log"
	"errors"
)

type ErrorService struct {
	err error
}

func (e *ErrorService) Error() string {
	return "Unknown Action Error"
}

func (e *ErrorService) UnknownActionError (messageIn *userConnections.MessageIn, chanOut chan <- *serviceModels.MessageOut) {
	e.err = errors.New(e.Error())
	messageOut := serviceModels.MessageOut{User:messageIn.User}
	messageOut.Err = e.err.Error()

	log.Println(messageOut.Err, " :ERROR SERVICE")
	chanOut <- &messageOut
}


