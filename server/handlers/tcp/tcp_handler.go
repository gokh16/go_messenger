package tcp

import (
	"bufio"
	"encoding/json"
	"fmt"
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

func Handler() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print("Connection doesn't accepted: ")
			log.Fatal(err)
		}
		go HandleJSON(conn)
	}
}

func HandleJSON(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from " + remoteAddr)
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Print("Data didn't read right: ")
			continue
		}
		ParseJSON([]byte(data), conn)
	}
}

func ParseJSON(bytes []byte, conn net.Conn) (Message, string, string) {
	flag := "tcp"
	message := Message{}
	err := json.Unmarshal(bytes, &message)
	if err != nil {
		log.Print("Unmarshal doesn't work: ")
		log.Fatal(err)
	}
	fmt.Println(message.Login)
	fmt.Println(message.Content)
	conn.Write([]byte(message.Content))
	return message, "func", flag
}
