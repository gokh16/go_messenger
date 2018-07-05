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
	go func() {
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
		DrawAuthWindow(conn)
	})
	if err != nil {
		log.Fatal(err)
	}
}
