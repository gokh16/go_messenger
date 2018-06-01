package userConnections

import (
	"github.com/gorilla/websocket"
	"net"
)

var WSConnections = map[*websocket.Conn]string{} // connection:login
var TCPConnections = map[net.Conn]string{}       // connection:login
