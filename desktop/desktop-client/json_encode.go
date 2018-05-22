package desktop_client

import (
	"encoding/json"
	"log"
	"fmt"
)
type User struct {
	Login    string
	//Password string
	//Username string `json:"username"`
	//Email    string `json:"email"`
	//Status   bool
	//UserIcon string
}

type Message struct {
	User               User
	//Group              Group
	//MessageContentType MessageContentType
	Content              string `json:"message_content"`
	//MessageSenderID      uint   `json:"message_sender_id"`
	//MessageRecipientID   uint   `json:"message_recepient_id"`
	//MessageContentTypeID uint   `json:"message_content_type_id"`
}

func JSONencode(user string, message string, function string) string{
	incomingData := Message{User{user}, message}
	outcomingData, err := json.Marshal(incomingData)
	if err!=nil{
		log.Fatal()
	}
	fmt.Println(string(outcomingData))
	fmt.Println(outcomingData)
	return string(outcomingData)+"\n"
}