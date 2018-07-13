package windows

import (
	"net"

	"github.com/ProtonMail/ui"
)

//DrawErrorWindow is a func which draw window by GTK's help
func DrawErrorWindow(errorText string, conn net.Conn) {
	window := ui.NewWindow("Error", 100, 50, false)
	errorLabel := ui.NewLabel(errorText)
	okButton := ui.NewButton("OK")
	mainBox := ui.NewVerticalBox()
	mainBox.Append(errorLabel, false)
	mainBox.Append(okButton, false)
	okButton.OnClicked(func(*ui.Button) {
		window.Destroy()
		if errorText == "You need to fill fields first!" {
			window.Hide()
			DrawRegistrationWindow(conn)
		}
		if errorText == "Enter the password!" {
			window.Hide()
			DrawAuthWindow(conn)
		}
		if errorText == "Wrong login or password" {
			window.Hide()
			DrawAuthWindow(conn)
		}
		if errorText == "Invalid email!" {
			window.Hide()
			DrawRegistrationWindow(conn)
		}
	})
	window.OnClosing(func(*ui.Window) bool {
		window.Destroy()
		return true
	})
	window.SetChild(mainBox)
	window.Show()
}
