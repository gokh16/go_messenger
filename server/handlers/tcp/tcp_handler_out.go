package tcp

import (
	"go_messenger/server/userConnections"
	"net"
	"encoding/json"
	"log"
	"fmt"
)

func WaitJSON(conns []net.Conn, str *userConnections.Message) {
	//for k,v := range conns {
	//	fmt.Println("TCP connect: ", k, v)
	//	fmt.Println(str.UserName)
	//}
	outcomingData, err := json.Marshal(str)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(conns)
	fmt.Println(outcomingData)
	for _, conn := range conns {
		conn.Write(outcomingData)
		conn.Write([]byte("\n"))
		fmt.Println("HERE")
	}
}

