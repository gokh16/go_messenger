package dekstopclient

import "github.com/andlabs/ui"

func Draw(){
	err := ui.Main(func() {
		window := ui.NewWindow("Chat", 500, 500, false)
		input := ui.NewEntry()
		send := ui.NewButton("Send")
		output := ui.NewLabel("")
		box := ui.NewVerticalBox()
		box.Append(output, true)
		box.Append(input, false)
		box.Append(send, false)

		send.OnClicked(func(*ui.Button) {
			output.SetText("Login -->  " + input.Text())
			input.SetText("")
		})
		window.SetChild(box)
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err!=nil{
		panic(err)
	}
}