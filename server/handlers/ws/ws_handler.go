package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go_messenger/server/userConnections"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WSHandler struct {
	Connection *userConnections.Connections
}

func (c *WSHandler) NewWSHandler(conns *userConnections.Connections) {
	ws := WSHandler{conns}
	Handler(ws)
}

func Handler(str WSHandler) {
	fs := http.FileServer(http.Dir("./public"))
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
		panic(err)
	}
}

func ReadMessage(conn *websocket.Conn, str WSHandler) {
	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			log.Println("Cannot read message")
		}
		GetJSON([]byte(data), conn, str)
	}
}

func GetJSON(bytes []byte, conn *websocket.Conn, str WSHandler) chan *userConnections.Message {
	message := userConnections.Message{}
	err := json.Unmarshal(bytes, &message)
	if err != nil {
		log.Println("Unmarshal error")
	}
	fmt.Println(message.UserName)
	fmt.Println(message.Content)
	str.Connection.AddWSConn(conn, message.UserName, &message)
	return str.Connection.OutChan
}

func SendJSON(conns []*websocket.Conn, str userConnections.Message) {
	outcomingData, err := json.Marshal(&str)
	if err != nil {
		log.Println(err)
	}
	for _, conn := range conns {
		conn.WriteJSON(outcomingData)
	}
}
