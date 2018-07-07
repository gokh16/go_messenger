package tcp

import (
	"encoding/json"
	"go_messenger/server/service/serviceModels"
	"log"
	"net"
)

//WaitJSON is waiting for data from route out, parsing data into json format and write to util(linter)
func WaitJSON(conns []net.Conn, str *serviceModels.MessageOut) {
	outComingData, err := json.Marshal(str)
	if err != nil {
		log.Println(err)
	}
	for _, conn := range conns {
		//todo ask how i may optimize it!
		if _, err := conn.Write(outComingData); err != nil {
			log.Println(err)
		}
		if _, err = conn.Write([]byte("\n")); err != nil {
			log.Println(err)
		}
	}
}
