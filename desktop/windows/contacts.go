package windows

import (
	"net"

	"go_messenger/desktop/structure"
	"log"

	"go_messenger/desktop/util"

	"go_messenger/desktop/config"

	"github.com/ProtonMail/ui"
)

//DrawContactsWindow is a func which draw window by GTK's help
func DrawContactsWindow(conn net.Conn, chatWindow *ui.Window) {
	log.Println("Opened DrawChatWindow")
	window := ui.NewWindow("Contacts", 400, 250, false)
	searchInput := ui.NewEntry()
	searchButton := ui.NewButton("Find")
	searchBox := ui.NewHorizontalBox()
	searchBox.Append(searchInput, true)
	searchBox.Append(searchButton, false)
	usersBox := ui.NewVerticalBox()
	mainBox := ui.NewVerticalBox()
	mainBox.Append(searchBox, false)
	users := make([]structure.User, 0)
	buttonUserSlice := make([]*ui.Button, 0)
	go func() {
		for {
			users = <-Contacts
			log.Println("Routine for read and show contacts")
			for _, user := range users {
				if user.Login != config.Login {
					buttonWithUser := ui.NewButton(user.Login)
					usersBox.Append(buttonWithUser, false)
					buttonUserSlice = append(buttonUserSlice, buttonWithUser)
					util.ContactsAction(buttonWithUser, conn, window, chatWindow)
				}
			}
			mainBox.Append(usersBox, true)
			break
		}
	}()
	window.SetChild(mainBox)
	window.OnClosing(func(*ui.Window) bool {
		window.Hide()
		users = nil
		return true
	})
	window.Show()
}
