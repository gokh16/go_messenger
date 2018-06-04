package tcp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go_messenger/server/userConnections"
	"log"
	"net"
)

type Message struct {
	UserName     string
	RelatingUser string
	RelatedUser  string
	RelationType uint
	GroupName    string
	GroupType    uint
	GroupOwner   string
	GroupMember  []string
	ContentType  uint
	Content      string
	LastMessage  string
	Login        string
	Password     string
	Email        string
	Status       bool
	UserIcon     string
	Action       string
}

var connections []net.Conn

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
			log.Fatal(err)
		}
		ParseJSON([]byte(data), conn)
	}
}

func ParseJSON(bytes []byte, conn net.Conn) (Message, string) {
	message := Message{}
	err := json.Unmarshal(bytes, &message)
	if err != nil {
		log.Print("Unmarshal doesn't work: ")
		log.Fatal(err)
	}
	fmt.Println(message.UserName)
	fmt.Println(message.Content)
	userConnections.TCPConnections[conn] = message.UserName
	for conns := range userConnections.TCPConnections {
		conns.Write([]byte(message.Content))
		conns.Write([]byte("\n"))
	}
	return message, " func "
}

func WaitJSON(conns []net.Conn, str Message) {
	outcomingData, err := json.Marshal(&str)
	if err != nil {
		log.Println(err)
	}
	for _, conn := range conns {
		conn.Write(outcomingData)
	}
}
