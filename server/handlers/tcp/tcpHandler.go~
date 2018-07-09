package tcp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go_messenger/server/routerIn"
	"go_messenger/server/userConnections"
	"log"
	"net"
)

//HandlerTCP is a structure which has attribute to connect with source structure in userConnections
type HandlerTCP struct {
	Connection *userConnections.Connections
}

//NewHandlerTCP is a constructor for TCP handler
func NewHandlerTCP(conns *userConnections.Connections) {
	tcp := HandlerTCP{conns}
	go tcp.Handler()
}

//Handler is a main func which is establish connections and call func for reading data from
//connection
func (t *HandlerTCP) Handler() {
	//todo ask about binds!
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := ln.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print("Connection doesn't accepted: ")
			log.Fatal(err)
		}

		go HandleJSON(conn, t)
	}
}

//HandleJSON method is handling json and call parser
func HandleJSON(conn net.Conn, str *HandlerTCP) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from " + remoteAddr)
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Printf("Client %v is gone!\n", str.Connection.GetUserNameByTCPConnection(conn))
			str.Connection.DeleteTCPConn(conn)
			log.Printf("ONLINE TCP CONNECTS AFTER DISCONNECT: -> %v", len(str.Connection.GetAllTCPConnections()))
			break
		}
		ParseJSON([]byte(data), conn, str)
	}
}

//ParseJSON method which advocates like parser
func ParseJSON(bytes []byte, conn net.Conn, str *HandlerTCP) {
	message := userConnections.MessageIn{}
	err := json.Unmarshal(bytes, &message)
	if err != nil {
		log.Print("Unmarshal doesn't work: ")
		log.Fatal(err)
	}
	log.Println(message.Group.GroupName, message.User.Username, message.Message.Content, "tcp_handler.go 72")
	str.Connection.AddTCPConn(conn, message.User.Username)
	routerIn.RouterIn(&message, str.Connection.OutChan)
}
