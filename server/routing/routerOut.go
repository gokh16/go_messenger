package routing

import (
	"github.com/gorilla/websocket"
	"go_messenger/server/userConnections"
	"net"
)

var message userConnections.Message

func RouterOut(msg userConnections.Message) {
	message = msg
	if listOfTCPcon := getAllTCPConnections(); listOfTCPcon != nil {
		//tcp.sendToTCP(listOfTCP, msg)
	}
	if listOfWScon := getAllWSConnections(); listOfWScon != nil {
		//ws.sendToWS(listOfWS, msg)
	}
}

func getAllTCPConnections() []net.Conn {
	mapTCP := userConnections.TCPConnections
	sliseTCPCon := []net.Conn{}
	for k, _ := range mapTCP {
		if mapTCP[k] == message.UserName /* depend on the Message structure on service level*/ {
			sliseTCPCon = append(sliseTCPCon, k)
		}
	}
	return sliseTCPCon
}

func getAllWSConnections() []*websocket.Conn {
	mapWS := userConnections.WSConnections
	sliseWSCon := []*websocket.Conn{}
	for k, _ := range mapWS {
		if mapWS[k] == message.UserName {
			sliseWSCon = append(sliseWSCon, k)
		}
	}
	return sliseWSCon
}
