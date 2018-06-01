package main

import (
	"./handlers/ws"
	"flag"
	"fmt"
	"go_messenger/server/handlers/tcp"
	"log"
	"net/http"
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
	go wsHandler()
	fmt.Println("good")
	go tcp.Handler()
	fmt.Println("good")

}
