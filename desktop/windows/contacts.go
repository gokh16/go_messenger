package windows

import (
	"net"

	"github.com/ProtonMail/ui"
)

//DrawContactsWindow is a func which draw window by GTK's help
func DrawContactsWindow(conn net.Conn, chatWindow *ui.Window) {
	//	log.Println("Opened DrawChatWindow")
	//	window := ui.NewWindow("Contacts", 400, 250, false)
	//	usersBox := ui.NewVerticalBox()
	//	users := make([]structure.User, 0)
	//	buttonUserSlice := make([]*ui.Button, 0)
	//	go func() {
	//		for {
	//			users = <- config.Contacts
	//			log.Println("Routine for read and show contacts (NEED TO BE FIXED)")
	//			log.Println(config.Contacts)
	//			for _, user := range users {
	//				buttonWithUser := ui.NewButton(user.Login)
	//				usersBox.Append(buttonWithUser, false)
	//				buttonUserSlice = append(buttonUserSlice, buttonWithUser)
	//			}
	//			for i := 0; i < len(buttonUserSlice); i++ {
	//				log.Println(i)
	//				util.ContactsAction(buttonUserSlice[i], conn, window, chatWindow)
	//			}
	//			window.SetChild(usersBox)
	//			break
	//		}
	//	}()
	//	//go func() {
	//	//	log.Println("Routine for check and redraw chat window (NEED TO BE FIXED)")
	//	//	for {
	//	//		window.Hide()
	//	//		chatWindow.Show()
	//	//		DrawChatWindow(conn)
	//	//	}
	//	//}()
	//
	//	window.OnClosing(func(*ui.Window) bool {
	//		window.Hide()
	//		users = nil
	//		return true
	//	})
	//	window.Show()
}
