package desktop_client

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
)

type Message struct {
	UserName    string
	GroupName   string
	ContentType string
	Content     string
	Login       string
	Password    string
	Email       string
	Status      bool
	UserIcon    string
	Action      string
}

func JSONencode(user string, groupName string, contentType string, content string, login string, password string, email string, status bool, userIcon string, action string) string {
	incomingData := Message{user, groupName, contentType, content, login, password, email, status, userIcon, action}
	outcomingData, err := json.Marshal(incomingData)
	if err != nil {
		log.Fatal(err)
	}
	return string(outcomingData) + "\n"
}

func JSONdecode(conn net.Conn) Message {
	message := Message{}
	jsonObj, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(jsonObj, &message)
	return message
}
