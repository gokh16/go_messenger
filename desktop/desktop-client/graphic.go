package desktop_client

import (
	"github.com/ProtonMail/ui"
	"net"
	"log"
	"fmt"
)


func Draw(){
	conn, _ := net.Dial("tcp", ":8080")
	err := ui.Main(func() {
		window := ui.NewWindow("Chat", 500, 500, false)
		input := ui.NewEntry()
		send := ui.NewButton("Send")
		output := ui.NewMultilineNonWrappingEntry()
		output.SetReadOnly(true)
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
			_, err := conn.Write([]byte(JSONencode(userExample1.Text(),input.Text(),"SendMessageTo")))
			if err!=nil{
				fmt.Println("OnClickedError!")
			}
			//writer,_ := io.Writer.Write(JSONencode(userExample1.Text(),input.Text(),"SendMessageTo"))
			fmt.Println(input.Text())
			//fmt.Fprintf(conn, input.Text()+"\n")
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
		log.Fatal(err)
	}
	//defer conn.Close()
}