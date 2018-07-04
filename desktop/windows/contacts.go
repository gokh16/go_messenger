package windows

import (
	"go_messenger/desktop/config"
	"go_messenger/desktop/structure"
	"go_messenger/desktop/util"
	"log"
	"net"
	"time"

	"github.com/ProtonMail/ui"
)

//DrawContactsWindow is a func which draw window by GTK's help
func DrawContactsWindow(conn net.Conn, chatWindow *ui.Window) {
	window := ui.NewWindow("Contacts", 400, 250, false)
	usersBox := ui.NewVerticalBox()
	channelForDrawUsers := make(chan []structure.User)
	users := make([]structure.User, 0)
	go func() {
		for {
			log.Println("wow")
			msg := util.JSONdecode(conn)
			users = msg.ContactList
			channelForDrawUsers <- users
			log.Println(users)
			break
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
			for i := 0; i < len(buttonUserSlice); i++ {
				log.Println(i)
				util.ContactsAction(buttonUserSlice[i], conn, window, chatWindow)
			}
			channelForDrawUsers <- nil
		}
	}()

	go func() {
		for {
			status := <-config.MarkForRedrawChatWindow
			if status == "groups are accepted" {
				window.Hide()
				chatWindow.Show()
				DrawChatWindow(conn)
			}
		}
	}()

	window.OnClosing(func(*ui.Window) bool {
		window.Hide()
		users = nil
		return true
	})
	window.Show()
}
