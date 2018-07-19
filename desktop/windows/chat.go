package windows

import (
	"go_messenger/desktop/config"
	"go_messenger/desktop/structure"
	"go_messenger/desktop/util"
	"log"
	"net"

	"github.com/ProtonMail/ui"
)

//DrawChatWindow is a func which draw window by GTK's help
func DrawChatWindow(conn net.Conn) {
	log.Println("Opened DrawChatWindow")
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
	refresh := ui.NewButton("Refresh")
	usersBox := ui.NewVerticalBox()
	usersBox.Append(refresh, false)

	buttonUserSlice := make([]*ui.Button, 0)
	status := make(chan bool)
	go func() {
		log.Println("Routine for accept, hang listeners and show data")
		json := <-Beginning
		log.Println(json.Status)
		if !json.Status {
			window.Hide()
			DrawErrorWindow("Wrong login or password", conn)
			status <- json.Status
			return
		}
		if json.Status {
			buttonUserSlice = nil
			for _, group := range config.UserGroups {
				if group != "" && group != config.Login {
					buttonWithGroup := ui.NewButton(group)
					usersBox.Append(buttonWithGroup, false)
					buttonUserSlice = append(buttonUserSlice, buttonWithGroup)
				}
			}
			for i := 0; i < len(buttonUserSlice); i++ {
				util.ButtonActions(buttonUserSlice[i], conn, output, json)
				output.SetText("")
			}
			status <- json.Status
			return
		}
	}()

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
		log.Println("Routine whis is printing input messages from server")
		for {
			json := <-Send
			log.Println(config.CurrentGroup, json.Message.MessageRecipientID)
			if config.CurrentGroup == json.Message.MessageRecipientID {
				output.Append(json.User.Login + ": " + json.Message.Content + "\n")
			}
		}
	}()
	refresh.OnClicked(func(*ui.Button) {
		window.Destroy()
		user := util.NewUser(config.Login, config.Password, config.Login, "test@test.com", true, "testUserIcon")
		message := util.NewMessageOut(user, &structure.User{}, &structure.Group{}, &structure.Message{}, nil, 1, 1, "LoginUser")
		_, err := conn.Write([]byte(util.JSONencode(*message)))
		if err != nil {
			log.Println("OnClickedError! Empty field.")
		}
		DrawChatWindow(conn)
	})
	send.OnClicked(func(*ui.Button) {
		log.Println("Button Send clicked")
		//FIX SLICEMEMBER
		output.Append(config.Login + ": " + input.Text() + "\n")
		id := config.GroupID[config.GroupName]
		//формирование новой структуры на отправку на сервер,
		//заполнение текущего экземпляра требуемыми полями.
		user := util.NewUser(config.Login, "", config.Login, "test@test.com", true, "testUserIcon")
		group := util.NewGroup(user, config.GroupName, config.UserID, 1)
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
		log.Println("Button Contacts clicked")
		user := util.NewUser(config.Login, "", config.Login, "test@test.com", true, "testUserIcon")
		message := util.NewMessageOut(user, &structure.User{}, &structure.Group{}, &structure.Message{}, nil, 1, 0, "GetContactList")
		_, err := conn.Write([]byte(util.JSONencode(*message)))
		if err != nil {
			log.Println("OnClickedError! Empty field.")
		}
		DrawContactsWindow(conn, window)
	})
	window.SetChild(mainBox)
	window.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	if a := <-status; a {
		window.Show()
	}
	return
}
