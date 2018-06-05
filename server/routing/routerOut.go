package routing

import (
	"github.com/gorilla/websocket"
	"go_messenger/server/userConnections"
	"net"
)

var message userConnections.Message

func RouterOut(msg userConnections.Message) {
	message = msg
	if listOfTCPCon := getAllTCPConnections(); listOfTCPCon != nil {
		//tcp.WaitJSON(listOfTCPCon, message) //  <- tcp.Message type, need userConnections.Message type in tcp package
	}
	if listOfWSCon := getAllWSConnections(); listOfWSCon != nil {
		//ws.sendToWS(listOfWSCon, message)
	}
}

func getAllTCPConnections() []net.Conn {
	mapTCP := userConnections.TCPConnections
	sliceTCPCon := []net.Conn{}
	for k, _ := range mapTCP {
		if mapTCP[k] == message.UserName {
			sliceTCPCon = append(sliceTCPCon, k)
		}
	}
	return sliceTCPCon
}

func getAllWSConnections() []*websocket.Conn {
	mapWS := userConnections.WSConnections
	sliceWSCon := []*websocket.Conn{}
	for k, _ := range mapWS {
		if mapWS[k] == message.UserName {
			sliceWSCon = append(sliceWSCon, k)
		}
	}
	return sliceWSCon
}
