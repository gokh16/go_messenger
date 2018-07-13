package routerOut

import (
	"go_messenger/server/handlers/tcp"
	"go_messenger/server/handlers/ws"
	"go_messenger/server/userConnections"
	"net"

	"go_messenger/server/service/serviceModels"

	"log"

	"github.com/gorilla/websocket"
)

//RouterOut is a structure which has attribute to connect with source structure in userConnections
type RouterOut struct {
	Connection *userConnections.Connections
}

//Function InitRouterOut is an init for routerOut struct
func InitRouterOut(conn *userConnections.Connections) {
	initRout := RouterOut{}
	initRout.Connection = conn
	go initRout.Handler()
}

// Handler is a main func which is establish connections and call func for reading data from
//connection
func (r *RouterOut) Handler() {
	//var msg is (*) pointer of serviceModels.MessageOut struct
	for msg := range r.Connection.OutChan {
		log.Println(msg.Action)
		if sliceTCPCon := r.getSliceOfTCP(msg); sliceTCPCon != nil {
			tcp.SendJSON(sliceTCPCon, msg)
		}
		if sliceWSCon := r.getSliceOfWS(msg); sliceWSCon != nil {
			ws.SendJSON(sliceWSCon, msg)
		}
	}
}

func (r *RouterOut) getSliceOfTCP(msg *serviceModels.MessageOut) []net.Conn {
	//get current TCP connections
	mapTCP := r.Connection.GetAllTCPConnections()
	log.Printf("ONLINE TCP connects -> %d", len(mapTCP))
	var sliceTCP []net.Conn

	if msg.Action == r.getAction(msg) { //LoginUser", "GetUsers", "GetGroupList", "GetGroup", "Error
		for conn, onlineUser := range mapTCP {
			if onlineUser == msg.User.Login {
				sliceTCP = append(sliceTCP, conn)
			}
		}
	} else {
		for conn, onlineUser := range mapTCP {
			for _, user := range msg.Recipients {
				if onlineUser == user.Login && onlineUser != msg.User.Login {
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
	log.Printf("ONLINE WS connects -> %d", len(mapWS))
	var sliceWS []*websocket.Conn

	if msg.Action == r.getAction(msg) { //LoginUser", "GetUsers", "GetGroupList", "GetGroup", "Error", "GetUser"
		for conn, onlineUser := range mapWS {
			if onlineUser == msg.User.Login {
				sliceWS = append(sliceWS, conn)
			}
		}
	} else {
		for conn, onlineUser := range mapWS {
			for _, user := range msg.Recipients {
				if onlineUser == user.Login && onlineUser != msg.User.Login {
					sliceWS = append(sliceWS, conn)
				}
			}
		}
	}
	return sliceWS
}

func (r *RouterOut) getAction(msg *serviceModels.MessageOut) string {
	switch msg.Action {
	case "LoginUser", "GetUsers", "GetGroupList", "GetGroup", "GetUser", "Error":
		return msg.Action
	default:
		return "No matches found"
	}
}
