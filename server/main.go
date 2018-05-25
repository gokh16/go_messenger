package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gokh16/go_messenger/server/handlers"
)

//"github.com/gokh16/go_messenger/server/orm"

// "github.com/jinzhu/gorm"
// _ "github.com/jinzhu/gorm/dialects/postgres"

func main() {

	// dbinfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	// 	orm.DB_HOST, orm.DB_PORT, orm.DB_USER, orm.DB_NAME, orm.DB_PASSWORD, orm.DB_SSLMODE)
	// fmt.Println(dbinfo)
	// db, err := gorm.Open("postgres", dbinfo)
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	////--------------------------
	// fs := http.FileServer(http.Dir("./public"))
	// http.Handle("/", fs)
	// http.HandleFunc("/ws", handlers.HandleWsConnections)

	// go handlers.HandleMessages()

	// log.Println("HTTP server started on :12345")
	// err := http.ListenAndServe(":12345", nil)
	// if err != nil {
	// 	panic(err)
	// }

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
