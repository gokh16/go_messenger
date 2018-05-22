package tcp

import (
	"net"
	"log"
	"encoding/json"
	"bufio"
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

func Handler(){
	ln, err := net.Listen("tcp", ":8080")
	if err!=nil{
		log.Fatal(err)
	}
	defer ln.Close()
	for{
		conn, err := ln.Accept()
		if err!=nil{
			log.Fatal(err)
		}
		go HandleJSON(conn)
	}
}

func HandleJSON(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from " + remoteAddr)
	for {
		data, err:=bufio.NewReader(conn).ReadString('\n')
		if err!=nil {
			log.Fatal(err)
		}
		ParseJSON([]byte(data)) //TODO create verse broadcast and rewrite scanner for scan by strings
	}
}

func ParseJSON(bytes []byte) (Message, string, string){
	flag := "tcp"
	message := Message{}
	err := json.Unmarshal(bytes, &message)
	if err != nil {
		fmt.Println(string(bytes))
		fmt.Println("here")
		log.Fatal(err)
	}
	fmt.Println(message.User.Login)
	fmt.Println(message.Content)
	return message, "func", flag
}
