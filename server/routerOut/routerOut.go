package routerOut

import (
	"fmt"
	"github.com/gorilla/websocket"
	"go_messenger/server/handlers/tcp"
	"go_messenger/server/handlers/ws"
	"go_messenger/server/userConnections"
	"net"
)

type RouterOut struct {
	Connection *userConnections.Connections
}

var msg *userConnections.Message

//var sliceTCPCon []net.Conn

func NewRouterOut(conn *userConnections.Connections) {
	newRout := RouterOut{}
	newRout.Connection = conn
	go newRout.HandleOut()
}

func (r *RouterOut) HandleOut() {
	//sliceTCP := r.getSliceOfTCP(r.Connection.OutChan)
	//sliceWS := r.getSliceOfWS(r.Connection.OutChan)
	//
	//for msg := range r.ChOut {
	//	if sliceTCP != nil {
	//		tcp.WaitJSON(sliceTCP, msg)
	//	}
	//	if sliceWS != nil {
	//		ws.SendJSON(sliceWS, msg)
	//	}
	//}

	for {
		if msg = <-r.Connection.OutChan; msg != nil {
			if sliceTCPCon := r.getSliceOfTCP(msg); sliceTCPCon != nil {
				tcp.WaitJSON(sliceTCPCon, msg)
			}
			if sliceWScon := r.getSliceOfWS(msg); sliceWScon != nil {
				ws.SendJSON(sliceWScon, msg)
			}
		}
	}
}

func (r *RouterOut) getSliceOfTCP(ms *userConnections.Message) []net.Conn {
	mapTCP := r.Connection.GetAllTCPConnections()
	fmt.Println("OLINE TCP connects -> ", len(mapTCP))
	var sliceTCP = []net.Conn{}
	for k, _ := range mapTCP {
		if mapTCP[k] == ms.UserName {
			sliceTCP = append(sliceTCP, k)
		}
	}
	return sliceTCP
}

func (r *RouterOut) getSliceOfWS(ms *userConnections.Message) []*websocket.Conn {
	mapWS := r.Connection.GetAllWSConnections()
	fmt.Println("OLINE WS connects -> ", len(mapWS))
	var sliceWS = []*websocket.Conn{}
	for k, _ := range mapWS {
		if mapWS[k] == ms.UserName {
			sliceWS = append(sliceWS, k)
		}
	}
	return sliceWS
}
