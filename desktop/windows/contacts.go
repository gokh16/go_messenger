package windows

import (
	"net"
	"github.com/ProtonMail/ui"
	"go_messenger/desktop/structure"
	"go_messenger/desktop/config"
	"go_messenger/desktop/util"
	"log"
	"time"
)

func DrawContactsWindow(conn net.Conn) {
	window := ui.NewWindow("Contacts", 400, 250, false)
	usersBox := ui.NewVerticalBox()
	channelForDrawUsers := make(chan []structure.User)
	users := make([]structure.User, 0)
	go func() {
		for {
			msg := util.JSONdecode(conn)
			users = msg.ContactList
			channelForDrawUsers <- users
			log.Println(users)
		}
	}()
	time.Sleep(1 * time.Second)
	buttonUserSlice := make([]*ui.Button, 0)
	go func() {
		if channelForDrawUsers != nil {
			for _, user := range users {
				log.Println(config.Login, user.Login)
				if user.Login != config.Login {
					buttonWithUser := ui.NewButton(user.Login)
					usersBox.Append(buttonWithUser, false)
					buttonUserSlice = append(buttonUserSlice, buttonWithUser)
				}
			}
			window.SetChild(usersBox)
			channelForDrawUsers <- nil
		}
	}()

	window.OnClosing(func(*ui.Window) bool {
		window.Hide()
		users = nil
		return true
	})
	window.Show()
}
