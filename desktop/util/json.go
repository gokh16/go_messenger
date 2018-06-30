package util

import (
	"bufio"
	"encoding/json"
	"go_messenger/desktop/structure"
	"log"
	"net"
)

//Message is a structure which has fields with data for send
//to the server
type MessageOut struct {
	User         structure.User
	Contact      structure.User
	Group        structure.Group
	Message      structure.Message
	Members      []structure.User
	RelationType uint
	MessageLimit uint
	Action       string
}

//MessageIn responce struct
type MessageIn struct {
	User        structure.User
	Members     []structure.User
	ContactList []structure.User
	GroupList   []Group
	Message     structure.Message
	Status      bool
	Action      string
	Err         string
}

type Group struct {
	ID        uint
	GroupName string
	GroupType structure.GroupType
	Members   []structure.User
	Messages  []structure.Message
}

//JSONencode is encoding source data to json
func JSONencode(message MessageOut) string {
	outcomingData, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}
	return string(outcomingData) + "\n"
}

//JSONdecode is decoding source json to message structure
func JSONdecode(conn net.Conn) MessageIn {
	message := MessageIn{}
	jsonObj, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}
	jsonError := json.Unmarshal(jsonObj, &message)
	if jsonError != nil {
		log.Println(jsonError, "unmarshaling")
	}
	if message.Err != "" {
		log.Println(message.Err,  " :CLIENT")
	}
	return message
}
