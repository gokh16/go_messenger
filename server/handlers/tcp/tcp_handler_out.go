package tcp

import (
	"net"
	"encoding/json"
	"log"
	"go_messenger/server/service/serviceModels"
)

func WaitJSON(conns []net.Conn, msg *serviceModels.MessageOut) {
	outComingData, err := json.Marshal(&msg)
	if err != nil {
		log.Println(err)
	}
	for _, conn := range conns {
		conn.Write(outComingData)
	}
}

