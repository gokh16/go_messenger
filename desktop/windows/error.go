package windows

import (
	"github.com/ProtonMail/ui"
)

func DrawErrorWindow(errorText string) {
	window := ui.NewWindow("Error", 100, 50, false)
	errorLabel := ui.NewLabel(errorText)
	okButton := ui.NewButton("OK")
	mainBox := ui.NewVerticalBox()
	mainBox.Append(errorLabel, false)
	mainBox.Append(okButton, false)
	okButton.OnClicked(func(*ui.Button) {
		window.Destroy()
	})
	window.OnClosing(func(*ui.Window) bool {
		window.Destroy()
		return true
	})
	window.SetChild(mainBox)
	window.Show()
}
