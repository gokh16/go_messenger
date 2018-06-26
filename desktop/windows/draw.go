package windows

import (
	"net"
	"log"
	"github.com/ProtonMail/ui"
)

//Draw func which must configure connection and draw window
//with further hierarchy
func Draw() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Println(err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
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