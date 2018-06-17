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

var msg *userConnections.Message

//NewRouterOut is a constructor for router out
func NewRouterOut(conn *userConnections.Connections) {
	newRout := RouterOut{}
	newRout.Connection = conn
	go newRout.Handler()
}

// Handler is a main func which is establish connections and call func for reading data from
////connection
func (r *RouterOut) Handler() {

	for {
		//TODO RANGE, IS OK!
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
		for k := range mapTCP {
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
		for k := range mapWS {
			if mapWS[k] == ms.GroupName {
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

//type RouterOut struct {
//	Connection *userConnections.Connections
//}
//
//var msg *userConnections.Message
//
////var sliceTCPCon []net.Conn
//
//func NewRouterOut(conn *userConnections.Connections) {
//	newRout := RouterOut{}
//	newRout.Connection = conn
//	go newRout.Handler()
//}
//
//func (r *RouterOut) Handler() {
//	//sliceTCP := r.getSliceOfTCP(r.Connection.OutChan)
//	//sliceWS := r.getSliceOfWS(r.Connection.OutChan)
//	//
//	//for msg := range r.ChOut {
//	//	if sliceTCP != nil {
//	//		tcp.WaitJSON(sliceTCP, msg)
//	//	}
//	//	if sliceWS != nil {
//	//		ws.SendJSON(sliceWS, msg)
//	//	}
//	//}
//
//	for {
//		if msg = <-r.Connection.OutChan; msg != nil {
//			if sliceTCPCon := r.getSliceOfTCP(msg); sliceTCPCon != nil {
//				tcp.WaitJSON(sliceTCPCon, msg)
//			}
//			if sliceWScon := r.getSliceOfWS(msg); sliceWScon != nil {
//				ws.SendJSON(sliceWScon, msg)
//			}
//		}
//	}
//}

//func (r *RouterOut) getSliceOfTCP(ms *userConnections.Message) []net.Conn {
//	mapTCP := r.Connection.GetAllTCPConnections()
//	fmt.Println("OLINE TCP connects -> ", len(mapTCP))
//	var sliceTCP = []net.Conn{}
//	for k, _ := range mapTCP {
//		//TEMP
//		if mapTCP[k] == ms.UserName {
//			sliceTCP = append(sliceTCP, k)
//		}
//	}
//	return sliceTCP
//}
//
//func (r *RouterOut) getSliceOfWS(ms *userConnections.Message) []*websocket.Conn {
//	mapWS := r.Connection.GetAllWSConnections()
//	fmt.Println("OLINE WS connects -> ", len(mapWS))
//	var sliceWS = []*websocket.Conn{}
//	for k, _ := range mapWS {
//		//TEMP
//		if mapWS[k] == ms.UserName {
//			sliceWS = append(sliceWS, k)
//		}
//	}
//	return sliceWS
//}
