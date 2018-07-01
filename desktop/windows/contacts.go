package windows

import (
	"net"
	"github.com/andlabs/ui"
	"go_messenger/desktop/structure"
	"go_messenger/desktop/config"
)

func DrawContactsWindow(conn net.Conn, users []*structure.User){
	window := ui.NewWindow("Contacts", 400,250, false)
	usersBox := ui.NewVerticalBox()
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
		window.Destroy()
		return true
	})
	window.Show()
}
