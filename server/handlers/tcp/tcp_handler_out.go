package tcp

import (
	"go_messenger/server/userConnections"
	"net"
	"encoding/json"
	"log"
)

func WaitJSON(conns []net.Conn, str *userConnections.Message) {
	outcomingData, err := json.Marshal(&str)
	if err != nil {
		log.Println(err)
	}
	for _, conn := range conns {
		conn.Write(outcomingData)
	}
}

