package routerOut

import (
	"github.com/gorilla/websocket"
	tcp2 "go_messenger/server/handlers/tcp"
	"go_messenger/server/userConnections"
	"net"
)

type RouterOut struct {
	Connection *userConnections.Connections
	ChOut      chan *userConnections.Message
}

func NewRouterOut(conn *userConnections.Connections, chOut chan *userConnections.Message) *RouterOut {
	newRout := RouterOut{}
	newRout.Connection = conn
	newRout.ChOut = chOut
	go newRout.HandleOut()
	return &newRout
}

func (r *RouterOut) HandleOut() {
	tcp := r.getSliceOfTCP(r.ChOut)
	ws := r.getSliceOfWS(r.ChOut)

	for msg := range r.ChOut {
		if tcp != nil {
			tcp2.WaitJSON(tcp, *msg)
		}
		if ws != nil {
			/*...*/
		}
	}
}

func (r *RouterOut) getSliceOfTCP(c chan *userConnections.Message) []net.Conn {
	mapTCP := r.Connection.GetAllTCPConnections()
	var sliceTCP = make([]net.Conn, len(mapTCP))
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
