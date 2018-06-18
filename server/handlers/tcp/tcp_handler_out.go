package tcp

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"go_messenger/server/service/serviceModels"
)

//WaitJSON is waiting for data from route out, parsing data into json format and write to client
func WaitJSON(conns []net.Conn, str *serviceModels.MessageOut) {
	outcomingData, err := json.Marshal(str)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(conns)
	fmt.Println(outcomingData)
	for _, conn := range conns {
		//todo ask how i may optimize it!
		_, err := conn.Write(outcomingData)
		if err != nil {
			log.Println(err)
		}
		_, err = conn.Write([]byte("\n"))
		if err != nil {
			log.Println(err)
		}
		fmt.Println("HERE")
	}
}
