package client

import (
	"encoding/json"
	"log"
	"go_messenger/desktop/structure"
)

//Message is a structure which has fields with data for send
//to the server
type MessageOut struct {
	User      structure.User
	Group     structure.Group
	GroupType structure.GroupType
	Message   structure.Message
	Action    string
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
//func JSONdecode(conn net.Conn) Message {
//	message := Message{}
//	jsonObj, err := bufio.NewReader(conn).ReadBytes('\n')
//	if err != nil {
//		log.Println(err)
//	}
//	jsonError := json.Unmarshal(jsonObj, &message)
//	if jsonError != nil {
//		log.Println(jsonError)
//	}
//	return message
//}
