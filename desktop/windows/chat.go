package windows

import (
	"net"
	"fmt"
	"github.com/ProtonMail/ui"
	"log"
	"go_messenger/desktop/structure"
	"go_messenger/desktop/util"
	"go_messenger/desktop/config"
	"time"
)

func DrawChatWindow(conn net.Conn) *ui.Window {
	time.Sleep(30 * time.Millisecond)
	window := ui.NewWindow(config.Login, 500, 500, false)
	input := ui.NewEntry()
	input.SetText("message")
	send := ui.NewButton("Send")
	output := ui.NewMultilineNonWrappingEntry()
	output.SetReadOnly(true)
	mainBox := ui.NewHorizontalBox()
	usersBox := ui.NewVerticalBox()
	buttonUserSlice := make([]*ui.Button, 0)
	for _, group := range config.UserGroups {
		if group != "" && group != config.Login {
			buttonWithUser := ui.NewButton(group)
			usersBox.Append(buttonWithUser, false)
			buttonUserSlice = append(buttonUserSlice, buttonWithUser)
		}
	}
	for i := 0; i < len(buttonUserSlice); i++ {
		util.ButtonActions(buttonUserSlice[i], conn, output)
		output.SetText("")
	}
	messageBox := ui.NewVerticalBox()
	messageBox.Append(output, true)
	messageBox.Append(input, false)
	messageBox.Append(send, false)
	mainBox.Append(usersBox, false)
	mainBox.Append(messageBox, true)
	go func() {
		status := <-config.MessagesGet
		if status {
			for {
				msg := util.JSONdecode(conn)
				if msg.Message.Content != "" {
					output.Append(msg.User.Login + ": " + msg.Message.Content + "\n")
				}
				log.Println(msg.Message.Content, "hjkhjkhjk")
				//todo подтягивать сообщение из базы
				//todo create update timeout

				fmt.Println(msg.Status)
			}
		}
	}()
	send.OnClicked(func(*ui.Button) {
		//FIX SLICEMEMBER
		output.Append(config.Login + ": " + input.Text() + "\n")
		id := config.GroupID[config.GroupName]
		//формирование новой структуры на отправку на сервер,
		//заполнение текущего экземпляра требуемыми полями.

		user:=util.NewUser(config.Login,"",config.Login, "test@test.com", true, "testUserIcon")
		group:=util.NewGroup(user,"private", config.GroupName, config.UserID, 1)
		msg := util.NewMessage(user, group, input.Text(), config.UserID, id,"Text")
		message := util.NewMessageOut(user, &structure.User{}, group, msg, nil, 1,0,"SendMessageTo")

		_, err := conn.Write([]byte(util.JSONencode(*message)))
		if err != nil {
			log.Println("OnClickedError! Empty field.")
		}
		input.SetText("")

	})
	window.SetChild(mainBox)
	window.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	window.Show()
	return window
}