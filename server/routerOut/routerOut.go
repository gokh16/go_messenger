package routerOut

import (
	"fmt"
	"go_messenger/server/handlers/tcp"
	"go_messenger/server/handlers/ws"
	"go_messenger/server/userConnections"
	"net"

	"github.com/gorilla/websocket"
)

//RouterOut is a structure which has attribute to connect with source structure in userConnections
type RouterOut struct {
	Connection *userConnections.Connections
}

//InitRouterOut is an init for router out
func InitRouterOut(conn *userConnections.Connections) {
	initRout := RouterOut{}
	initRout.Connection = conn
	go initRout.Handler()
}

// Handler is a main func which is establish connections and call func for reading data from
//connection
func (r *RouterOut) Handler() {

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
	fmt.Println("ONLINE TCP connects -> ", len(mapTCP))
	var sliceTCP []net.Conn

	if msg.Action == "GetUsers" {
		for conn, onlineUser := range mapTCP {
			if onlineUser == msg.UserName {
				sliceTCP = append(sliceTCP, conn)
			}
		}
	}

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

	if msg.Action == "GetUsers" {
		for conn, onlineUser := range mapWS {
			if onlineUser == msg.UserName {
				sliceWS = append(sliceWS, conn)
			}
		}
	}

	for conn, onlineUser := range mapWS {
		for _, groupMember := range msg.GroupMember {
			if onlineUser == groupMember && onlineUser != msg.UserName {
				sliceWS = append(sliceWS, conn)
			}
		}
	}
	return sliceWS
}
