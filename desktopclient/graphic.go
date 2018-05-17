package desktopclient

import "github.com/andlabs/ui"

func Draw(){
	err := ui.Main(func() {
		window := ui.NewWindow("Chat", 500, 500, false)
		input := ui.NewEntry()
		send := ui.NewButton("Send")
		output := ui.NewLabel("")
		mainBox := ui.NewHorizontalBox()
		userExample1 := ui.NewButton("USER1")
		userExample2 := ui.NewButton("USER2")
		userExample3 := ui.NewButton("USER3")
		userExample4 := ui.NewButton("USER4")
		usersBox := ui.NewVerticalBox()
		usersBox.Append(userExample1, false)
		usersBox.Append(userExample2, false)
		usersBox.Append(userExample3, false)
		usersBox.Append(userExample4, false)
		messageBox:=ui.NewVerticalBox()
		messageBox.Append(output, true)
		messageBox.Append(input, false)
		messageBox.Append(send, false)
		mainBox.Append(usersBox, false)
		mainBox.Append(messageBox, true)

		send.OnClicked(func(*ui.Button) {
			output.SetText("Login -->  " + input.Text())
			input.SetText("")
		})
		window.SetChild(mainBox)
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