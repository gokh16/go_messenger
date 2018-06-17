package client

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
)

//Message is a structure which has fields with data for send
//to the server
type Message struct {
	UserName     string
	RelatingUser string
	RelatedUser  string
	RelationType uint
	GroupName    string
	GroupType    uint
	GroupOwner   string
	GroupMember  []string
	ContentType  string
	Content      string
	LastMessage  string
	Login        string
	Password     string
	Email        string
	Status       bool
	UserIcon     string
	Action       string
}

//JSONencode is encoding source data to json
func JSONencode(user string, relatingUser string, relatedUser string, relationType uint, groupName string, groupType uint,
	groupOwner string, groupMember []string, contentType string, content string, lastMessage string, login string, password string,
	email string, status bool, userIcon string, action string) string {
	incomingData := Message{user, relatingUser, relatedUser, relationType, groupName, groupType, groupOwner,
		groupMember, contentType, content, lastMessage, login, password, email, status, userIcon, action}
	outcomingData, err := json.Marshal(incomingData)
	if err != nil {
		log.Fatal(err)
	}
	return string(outcomingData) + "\n"
}

//JSONdecode is decoding source json to message structure
func JSONdecode(conn net.Conn) Message {
	message := Message{}
	jsonObj, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		log.Println(err)
	}
	jsonError := json.Unmarshal(jsonObj, &message)
	if jsonError != nil {
		log.Println(jsonError)
	}
	return message
}
