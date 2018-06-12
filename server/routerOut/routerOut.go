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

func NewRouterOut(conn *userConnections.Connections) {
	newRout := RouterOut{}
	newRout.Connection = conn
	go newRout.HandleOut()
}

func (r *RouterOut) HandleOut() {

	for {
		if msg = <-r.Connection.OutChan; msg != nil {
			if sliceTCPCon := r.getSliceOfTCP(msg); sliceTCPCon != nil {
				tcp.WaitJSON(sliceTCPCon, msg)
			}
			if sliceWSCon := r.getSliceOfWS(msg); sliceWSCon != nil {
				ws.SendJSON(sliceWSCon, msg)
			}
		}
	}
}

func (r *RouterOut) getSliceOfTCP(ms *userConnections.Message) []net.Conn {

	//get current TCP connections
	mapTCP := r.Connection.GetAllTCPConnections()
	fmt.Println("ONLINE TCP connects -> ", len(mapTCP))
	var sliceTCP = []net.Conn{}

	//send message to the client
	fmt.Println(ms.GroupMember, "groupmember")
	if len(ms.GroupMember) == 0 {
		for k, _ := range mapTCP {
			if mapTCP[k] == ms.UserName {
				sliceTCP = append(sliceTCP, k)
			}
		}

		//send message to the group
	} else if len(ms.GroupMember) > 0 {
		for _, groupMember := range ms.GroupMember {
			for conn, onlineUserName := range mapTCP {
				if onlineUserName == groupMember {
					sliceTCP = append(sliceTCP, conn)
				}
			}
		}
	}
	return sliceTCP
}

func (r *RouterOut) getSliceOfWS(ms *userConnections.Message) []*websocket.Conn {

	//get current WS connections
	mapWS := r.Connection.GetAllWSConnections()
	fmt.Println("ONLINE WS connects -> ", len(mapWS))
	var sliceWS = []*websocket.Conn{}

	//send message to the client
	if len(ms.GroupMember) == 0 {
		for k, _ := range mapWS {
			if mapWS[k] == ms.UserName {
				sliceWS = append(sliceWS, k)
			}
		}

		//send message to the group
	} else if len(ms.GroupMember) > 0 {
		for _, groupMember := range ms.GroupMember {
			for conn, onlineUserName := range mapWS {
				if onlineUserName == groupMember {
					sliceWS = append(sliceWS, conn)
				}
			}
		}
	}
	return sliceWS
}
