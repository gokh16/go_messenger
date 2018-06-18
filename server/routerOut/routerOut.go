package routerOut

import (
	"fmt"
	"go_messenger/server/handlers/tcp"
	"go_messenger/server/handlers/ws"
	"go_messenger/server/service/serviceModels"
	"go_messenger/server/userConnections"
	"net"

	"github.com/gorilla/websocket"
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

	//var msg is (*) pointer of userConnections.MessageIn struct
	for msg := range r.Connection.OutChan {
		if sliceTCPCon := r.getSliceOfTCP(msg); sliceTCPCon != nil {
			tcp.WaitJSON(sliceTCPCon, msg)
		}
		if sliceWSCon := r.getSliceOfWS(msg); sliceWSCon != nil {
			ws.SendJSON(sliceWSCon, msg)
		}
	}
}

func (r *RouterOut) getSliceOfTCP(msg *serviceModels.MessageOut) []net.Conn {

	//get current TCP connections
	mapTCP := r.Connection.GetAllTCPConnections()
	fmt.Println("ONLINE TCP connects -> ", len(mapTCP))
	var sliceTCP []net.Conn

	if msg.Action == "GetUsers" {
		for conn, onlineUser := range mapTCP {
			if onlineUser == msg.User.Username {
				sliceTCP = append(sliceTCP, conn)
			}
		}
	}

	for conn, onlineUser := range mapTCP {
		for _, group := range msg.GroupList {
			for _, groupMember := range group.Members {
				if onlineUser == groupMember.Username && onlineUser != msg.User.Username{
					sliceTCP = append(sliceTCP, conn)
				}
			}
		}
	}
	return sliceTCP
}

func (r *RouterOut) getSliceOfWS(msg *serviceModels.MessageOut) []*websocket.Conn {

	//get current WS connections
	mapWS := r.Connection.GetAllWSConnections()
	fmt.Println("ONLINE WS connects -> ", len(mapWS))
	var sliceWS []*websocket.Conn

	if msg.Action == "GetUsers" {
		for conn, onlineUser := range mapWS {
			if onlineUser == msg.User.Username {
				sliceWS = append(sliceWS, conn)
			}
		}
	}

	for conn, onlineUser := range mapWS {
		for _, group := range msg.GroupList {
			for _, groupMember := range group.Members {
				if onlineUser == groupMember.Username && onlineUser != msg.User.Username {
					sliceWS = append(sliceWS, conn)
				}
			}
		}
	}
	return sliceWS
}
