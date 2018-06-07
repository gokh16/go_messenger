package routerOut

import (
	"github.com/gorilla/websocket"
	"go_messenger/server/userConnections"
	"net"
	"fmt"
	"go_messenger/server/handlers/tcp"
)

type RouterOut struct {
	Connection *userConnections.Connections
}

func NewRouterOut(conn *userConnections.Connections) *RouterOut {
	newRout := RouterOut{conn}
	go newRout.HandleOut()
	return &newRout
}

func (r *RouterOut) HandleOut() {
	//for {
	//	sliceTCP := r.getSliceOfTCP(r.Connection.OutChan)
	//	sliceWS := r.getSliceOfWS(r.Connection.OutChan)
	//	fmt.Println(sliceTCP)
	//	fmt.Println(sliceWS)
	//if sliceTCP != nil {
	//	msg := <-r.Connection.OutChan
	//	tcp.WaitJSON(sliceTCP, msg)
	//}
	//if sliceWS != nil {
	//	msg := <-r.Connection.OutChan
	//	ws.SendJSON(sliceWS, msg)
	//}
	//}
	for {
		if msg := <-r.Connection.OutChan; msg != nil {
			if sliceTCPCon := r.getSliceOfTCP(r.Connection.OutChan); sliceTCPCon != nil {
				tcp.WaitJSON(sliceTCPCon, msg)
			}
		}

	}
}

func (r *RouterOut) getSliceOfTCP(c chan *userConnections.Message) []net.Conn {
	mapTCP := r.Connection.GetAllTCPConnections()
	var sliceTCP = make([]net.Conn, len(mapTCP))
	fmt.Println(mapTCP)
	for m := range c {
		for k, _ := range mapTCP {
			if mapTCP[k] == m.UserName {
				sliceTCP = append(sliceTCP, k)
			}
		}
	}
	return sliceTCP
}

func (r *RouterOut) getSliceOfWS(c chan *userConnections.Message) []*websocket.Conn {
	mapWS := r.Connection.GetAllWSConnections()
	var sliceWS = make([]*websocket.Conn, len(mapWS))
	for m := range c {
		for k, _ := range mapWS {
			if mapWS[k] == m.UserName {
				sliceWS = append(sliceWS, k)
			}
		}
	}
	return sliceWS
}
