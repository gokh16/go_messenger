package windows

import (
	"log"
	"net"

	"github.com/ProtonMail/ui"
)

//Draw func which must configure connection and draw window
//with further hierarchy
func Draw() {
	conn, err := net.Dial("tcp", ":8080")
	log.Println("Creating connection")
	go func() {
		log.Println("Routine which is initializating reader")
		Reader(conn)
	}()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		er := conn.Close()
		if er != nil {
			log.Println(err)
		}
	}()
	err = ui.Main(func() {
		log.Println("Main func which create and draw the first window!")
		DrawAuthWindow(conn)
	})
	if err != nil {
		log.Fatal(err)
	}
}
