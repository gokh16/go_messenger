package tcp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go_messenger/server/userConnections"
	"log"
	"net"
)

type TCPHandler struct {
	Connection *userConnections.Connections
}

func (c *TCPHandler) NewTCPHandler (conns *userConnections.Connections) *TCPHandler{
	tcp := TCPHandler{conns}
	Handler()
	return &tcp
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
			log.Fatal(err)
		}
		ParseJSON([]byte(data), conn)
	}
}

func ParseJSON(bytes []byte, conn net.Conn) chan *userConnections.Message{
	message := userConnections.Message{}
	err := json.Unmarshal(bytes, &message)
	if err != nil {
		log.Print("Unmarshal doesn't work: ")
		log.Fatal(err)
	}
	fmt.Println(message.UserName)
	fmt.Println(message.Content)
	TCPHandler{}.Connection.AddTCPConn(conn,message.UserName,&message)
	return TCPHandler{}.Connection.OutChan
}