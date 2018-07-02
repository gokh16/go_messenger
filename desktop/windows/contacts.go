package windows

import (
	"net"
	"github.com/ProtonMail/ui"
	"go_messenger/desktop/structure"
	"go_messenger/desktop/config"
	"go_messenger/desktop/util"
	"log"
)

func DrawContactsWindow(conn net.Conn) {
	window := ui.NewWindow("Contacts", 400, 250, false)
	usersBox := ui.NewVerticalBox()
	users := make([]structure.User, 0)
	go func() {
		for {
			log.Println("contacts")
			msg := util.JSONdecode(conn)
			users = msg.ContactList
			log.Println(msg.Action, users[0].Login, "contacts")
		}
	}()
	buttonUserSlice := make([]*ui.Button, 0)
	for _, user := range users {
		if user.Login != config.Login {
			buttonWithUser := ui.NewButton(user.Login)
			usersBox.Append(buttonWithUser, false)
			buttonUserSlice = append(buttonUserSlice, buttonWithUser)
		}
	}
	window.SetChild(usersBox)
	window.OnClosing(func(*ui.Window) bool {
		window.Hide()
		return true
	})
	window.Show()
}
