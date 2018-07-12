package windows

import (
	"net"

	"go_messenger/desktop/structure"
	"log"

	"go_messenger/desktop/util"

	"github.com/ProtonMail/ui"
)

//DrawContactsWindow is a func which draw window by GTK's help
func DrawContactsWindow(conn net.Conn, chatWindow *ui.Window) {
	log.Println("Opened DrawChatWindow")
	window := ui.NewWindow("Contacts", 400, 250, false)
	users := make([]structure.User, 0)
	usersBox := ui.NewVerticalBox()
	buttonUserSlice := make([]*ui.Button, 0)
	go func() {
		for {
			users = <-Contacts
			log.Println("Routine for read and show contacts (NEED TO BE FIXED)")
			for _, user := range users {
				buttonWithUser := ui.NewButton(user.Login)
				usersBox.Append(buttonWithUser, false)
				buttonUserSlice = append(buttonUserSlice, buttonWithUser)
				util.ContactsAction(buttonWithUser, conn, window, chatWindow)
			}
			window.SetChild(usersBox)
			break
		}
	}()
	window.OnClosing(func(*ui.Window) bool {
		window.Hide()
		users = nil
		return true
	})
	window.Show()
}
