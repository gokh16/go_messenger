package routerOut

import (
	"fmt"
	"github.com/gorilla/websocket"
	"go_messenger/server/handlers/tcp"
	"go_messenger/server/handlers/ws"
	"go_messenger/server/userConnections"
	"log"
	"net"
)

type RouterOut struct {
	Connection *userConnections.Connections
}

func InitRouterOut(conn *userConnections.Connections) {
	initRout := RouterOut{}
	initRout.Connection = conn
	go initRout.HandleOut()
}

func (r *RouterOut) HandleOut() {

	//var msg is (*) pointer of userConnections.Message struct
	for msg := range r.Connection.OutChan {
		if sliceTCPCon := r.getSliceOfTCP(msg); sliceTCPCon != nil {
			tcp.WaitJSON(sliceTCPCon, msg)
		}
		if sliceWSCon := r.getSliceOfWS(msg); sliceWSCon != nil {
			ws.SendJSON(sliceWSCon, msg)
		}
	}
}

func (r *RouterOut) getSliceOfTCP(msg *userConnections.Message) []net.Conn {

	//get current TCP connections
	mapTCP := r.Connection.GetAllTCPConnections()
	log.Println("ONLINE TCP connects -> ", len(mapTCP))
	var sliceTCP []net.Conn

	for conn, onlineUser := range mapTCP {
		for _, groupMember := range msg.GroupMember {
			if onlineUser == groupMember && onlineUser != msg.UserName {
				sliceTCP = append(sliceTCP, conn)
			}
		}
	}
	return sliceTCP
}

func (r *RouterOut) getSliceOfWS(msg *userConnections.Message) []*websocket.Conn {

	//get current WS connections
	mapWS := r.Connection.GetAllWSConnections()
	fmt.Println("ONLINE WS connects -> ", len(mapWS))
	var sliceWS []*websocket.Conn

	for conn, onlineUser := range mapWS {
		for _, groupMember := range msg.GroupMember {
			if onlineUser == groupMember && onlineUser != msg.UserName {
				sliceWS = append(sliceWS, conn)
			}
		}
	}
	return sliceWS
}
