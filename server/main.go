package main

import (
	"./handlers/tcp"
	"flag"
	"log"
	"net/http"
	"fmt"
	"./handlers/ws"
)

func tcpHandler(){
	tcp.Handler()
}

func wsHandler(){
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
	go wsHandler()	
	fmt.Println("good")
	tcpHandler()
	fmt.Println("good")

}