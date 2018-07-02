package windows

import (
	"fmt"
	"go_messenger/desktop/config"
	"go_messenger/desktop/structure"
	"go_messenger/desktop/util"
	"log"
	"net"
	"time"

	"github.com/ProtonMail/ui"
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
		for {
			msg := util.JSONdecode(conn)
			if msg.Message.Content != "" {
				output.Append(msg.User.Login + ": " + msg.Message.Content + "\n")
			}
			log.Println(msg.Message.Content)
			//todo подтягивать сообщение из базы
			//todo create update timeout

			fmt.Println(msg.Status)
		}
	}()
	send.OnClicked(func(*ui.Button) {
		//FIX SLICEMEMBER
		output.Append(config.Login + ": " + input.Text() + "\n")
		id := config.GroupID[config.GroupName]
		log.Println(id, config.GroupName)
		//формирование новой структуры на отправку на сервер,
		//заполнение текущего экземпляра требуемыми полями.
		message := util.MessageOut{
			User: structure.User{
				Login:    config.Login,
				Password: "testPassword",
				Username: config.Login,
				Email:    "test@test.com",
				Status:   true,
				UserIcon: "testUserIcon",
			},
			Contact: structure.User{},
			Group: structure.Group{
				User: structure.User{
					Login:    config.Login,
					Password: "testPassword",
					Username: config.Login,
					Email:    "test@test.com",
					Status:   true,
					UserIcon: "testUserIcon",
				},
				GroupType: structure.GroupType{
					Type: "private",
				},
				GroupName: config.GroupName,
				//GroupOwnerID: config.UserID,
				GroupTypeID: 1,
			},
			Message: structure.Message{
				User: structure.User{
					Login:    config.Login,
					Password: "testPassword",
					Username: config.Login,
					Email:    "test@test.com",
					Status:   true,
					UserIcon: "testUserIcon",
				},
				MessageSenderID:    config.UserID,
				MessageRecipientID: id,
				Group: structure.Group{
					User: structure.User{
						Login:    config.Login,
						Password: "testPassword",
						Username: config.Login,
						Email:    "test@test.com",
						Status:   true,
						UserIcon: "testUserIcon",
					},
					GroupType: structure.GroupType{
						Type: "private",
					},
					GroupName: config.GroupName,
					//GroupOwnerID: 123,
					GroupTypeID: 1,
				},
				Content: input.Text(),
			},
			Members:      nil,
			RelationType: 1,
			MessageLimit: 1,
			Action:       "SendMessageTo",
		}
		log.Println(message.Group.GroupName)
		_, err := conn.Write([]byte(util.JSONencode(message)))
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
