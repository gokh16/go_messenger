package ws

import (
	"encoding/json"
	"fmt"
	"go_messenger/server/routerIn"
	"go_messenger/server/userConnections"
	"log"
	"net/http"

	"go_messenger/server/service/serviceModels"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//HandlerWS is a structure which has attribute to connect with source structure in userConnections
type HandlerWS struct {
	Connection *userConnections.Connections
}

//NewHandlerWS is a constructor for WS handler
func NewHandlerWS(conns *userConnections.Connections) {
	ws := HandlerWS{conns}
	go Handler(ws)
}

//Handler is a main func which is establish connections and call func for reading data from
//connection
func Handler(str HandlerWS) {
	fs := http.FileServer(http.Dir("../web"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Cannot upgrade")
		}
		go ReadMessage(conn, str)
	})
	log.Println("HTTP server started on :12345")
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Println(err)
	}
}

//ReadMessage is a func for reading data from ws connection
func ReadMessage(conn *websocket.Conn, str HandlerWS) {
	for {
		messageType, data, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Client %v is gone!\n", str.Connection.GetUserNameByWSConnection(conn))
			str.Connection.DeleteWSConnection(conn)
			log.Printf("ONLINE WS CONNECTS AFTER DISCONNECT: -> %v", len(str.Connection.GetAllWSConnections()))
			return
		}
		if err := conn.WriteMessage(messageType, data); err != nil {
			log.Println(err)
			return
		}
		GetJSON(data, conn, str)
		/*_, data, err := conn.ReadMessage()
		if err != nil {
			log.Println("Cannot read message")
		}*/
		//GetJSON(data, conn, str)
	}
}

//GetJSON is
func GetJSON(bytes []byte, conn *websocket.Conn, str HandlerWS) {
	message := userConnections.MessageIn{}
	err := json.Unmarshal(bytes, &message)
	if err != nil {
		log.Println("Unmarshal error")
	}
	str.Connection.AddWSConn(conn, message.User.Username)
	fmt.Println("gn", message.Group.GroupName)
	routerIn.RouterIn(&message, str.Connection.OutChan)
	//return str.Connection.outChan
}

//SendJSON is waiting for data from route out, parsing data into json format and write to util
func SendJSON(conns []*websocket.Conn, str *serviceModels.MessageOut) {
	fmt.Println("send", str.Message.Content)
	for _, conn := range conns {
		err := conn.WriteJSON(str)
		if err != nil {
			log.Println(err)
		}
	}
}
