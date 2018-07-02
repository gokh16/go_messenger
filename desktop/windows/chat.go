package windows

import (
	"go_messenger/desktop/config"
	"go_messenger/desktop/structure"
	"go_messenger/desktop/util"
	"log"
	"net"
	"time"

	"github.com/ProtonMail/ui"

	"github.com/ProtonMail/ui"
)

func DrawChatWindow(conn net.Conn) *ui.Window {
	time.Sleep(30 * time.Millisecond)
	window := ui.NewWindow(config.Login, 800, 500, false)
	input := ui.NewEntry()
	input.SetText("message")
	send := ui.NewButton("Send")
	profile := ui.NewButton("Profile")
	contacts := ui.NewButton("Contacts")
	userHeader := ui.NewHorizontalBox()
	output := ui.NewMultilineNonWrappingEntry()
	output.SetReadOnly(true)
	mainBox := ui.NewHorizontalBox()
	searchEntry := ui.NewEntry()
	searchEntry.SetText("Search")
	usersBox := ui.NewVerticalBox()
	usersBox.Append(searchEntry, false)
	buttonUserSlice := make([]*ui.Button, 0)
	for _, group := range config.UserGroups {
		if group != "" && group != config.Login {
			buttonWithGroup := ui.NewButton(group)
			usersBox.Append(buttonWithGroup, false)
			buttonUserSlice = append(buttonUserSlice, buttonWithGroup)
		}
	}
	for i := 0; i < len(buttonUserSlice); i++ {
		util.ButtonActions(buttonUserSlice[i], conn, output)
		output.SetText("")
	}
	userHeader.Append(profile, true)
	userHeader.Append(contacts, true)
	messageBox := ui.NewVerticalBox()
	messageBox.Append(userHeader, false)
	messageBox.Append(output, true)
	messageBox.Append(input, false)
	messageBox.Append(send, false)
	mainBox.Append(usersBox, false)
	mainBox.Append(messageBox, true)
	go func() {
		status := <-config.MarkForRead
		for {
			if status { //todo finish THIS PART!
				msg := util.JSONdecode(conn)
				if msg.Message.Content != "" {
					output.Append(msg.User.Login + ": " + msg.Message.Content + "\n")
				}
				log.Println(msg.Action, "chat window")
				//todo подтягивать сообщение из базы
				//todo create update timeout

			}
		}
	}()
	send.OnClicked(func(*ui.Button) {
		//FIX SLICEMEMBER
		output.Append(config.Login + ": " + input.Text() + "\n")
		id := config.GroupID[config.GroupName]
		//формирование новой структуры на отправку на сервер,
		//заполнение текущего экземпляра требуемыми полями.

		user := util.NewUser(config.Login, "", config.Login, "test@test.com", true, "testUserIcon")
		group := util.NewGroup(user, "private", config.GroupName, config.UserID, 1)
		msg := util.NewMessage(user, group, input.Text(), config.UserID, id, "Text")
		message := util.NewMessageOut(user, &structure.User{}, group, msg, nil, 1, 0, "SendMessageTo")
		if msg.Content != "" {
			_, err := conn.Write([]byte(util.JSONencode(*message)))
			if err != nil {
				log.Println("OnClickedError! Empty field.")
			}
			input.SetText("")
		}
	})
	contacts.OnClicked(func(*ui.Button) {
		user := util.NewUser(config.Login, "", config.Login, "test@test.com", true, "testUserIcon")
		message := util.NewMessageOut(user, &structure.User{}, &structure.Group{}, &structure.Message{}, nil, 1, 0, "GetUsers")
		_, err := conn.Write([]byte(util.JSONencode(*message)))
		if err != nil {
			log.Println("OnClickedError! Empty field.")
		}
		DrawContactsWindow(conn)
	})
	window.SetChild(mainBox)
	window.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	window.Show()
	return window
}
