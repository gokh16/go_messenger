package tcp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go_messenger/server/userConnections"
	"log"
	"net"
	"go_messenger/server/routerIn"
)

type TCPHandler struct {
	Connection *userConnections.Connections
}


func NewTCPHandler (conns *userConnections.Connections) {
	tcp := TCPHandler{conns}
	go tcp.Handler()
}

func (t *TCPHandler) Handler() {
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

		go HandleJSON(conn, t)
	}
}

func HandleJSON(conn net.Conn, str *TCPHandler) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from " + remoteAddr)
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Print("Data didn't read right: ")
			log.Fatal(err)
		}
		ParseJSON([]byte(data), conn, str)
	}
}


func ParseJSON(bytes []byte, conn net.Conn, str *TCPHandler) {
	message := userConnections.MessageIn{}
	err := json.Unmarshal(bytes, &message)
	if err != nil {
		log.Print("Unmarshal doesn't work: ")
		log.Fatal(err)
	}
	str.Connection.AddTCPConn(conn, message.User.Username)
	routerIn.RouterIn(&message, str.Connection.OutChan)
}