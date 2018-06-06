package main

import (
	"./handlers/ws"
	"flag"
	"fmt"
	"log"
	"net/http"
	"go_messenger/server/handlers/tcp"
	"go_messenger/server/userConnections"
)

func wsHandler() {
	flag.Parse()
	hub := handlers.NewHub()

	go hub.RunHub()
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeWebsocket(hub, w, r)
	})

	log.Println("HTTP server started on :12345")
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	connectionList := userConnections.Connections{}
	go wsHandler()
	fmt.Println("good")
	tcpStr := &tcp.TCPHandler{}
	tcpStr.NewTCPHandler(&connectionList)
	fmt.Println("good")

}
