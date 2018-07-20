package tcp

import (
	"encoding/json"
	"go_messenger/server/service/serviceModels"
	"log"
	"net"
)

//SendJSON is waiting for data from route out, parsing data into json format and write to util(linter)
func SendJSON(conns []net.Conn, str *serviceModels.MessageOut) {
	outComingData, err := json.Marshal(str)
	if err != nil {
		log.Println(err)
	}
	for _, conn := range conns {
		//todo ask how i may optimize it!
		if _, er := conn.Write(outComingData); er != nil {
			log.Println(err)
		}
		if _, er := conn.Write([]byte("\n")); er != nil {
			log.Println(err)
		}
	}
}
