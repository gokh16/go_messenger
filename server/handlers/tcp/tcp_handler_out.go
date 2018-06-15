package tcp

import (
	"go_messenger/server/userConnections"
	"net"
	"encoding/json"
	"log"
)

func WaitJSON(conns []net.Conn, msg *userConnections.Message) {
	outComingData, err := json.Marshal(&msg)
	if err != nil {
		log.Println(err)
	}
	for _, conn := range conns {
		conn.Write(outComingData)
	}
}

