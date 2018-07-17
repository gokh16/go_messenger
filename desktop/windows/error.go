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
		DrawAuthWindow(conn)
	})
	window.OnClosing(func(*ui.Window) bool {
		window.Destroy()
		return true
	})
	window.SetChild(mainBox)
	window.Show()
}
