package desktop_client

import (
	"net"
	"bufio"
	"log"
)

func Read(conn net.Conn) string{
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err!=nil{
		log.Fatal(err)
	}
	return string(message)
}