package util

import (
	"log"
	"net"

	"go_messenger/desktop/structure"

	"go_messenger/desktop/config"

	"github.com/ProtonMail/ui"
)

//ButtonActions is hanging listeners for group buttons
func ButtonActions(button *ui.Button, conn net.Conn, output *ui.MultilineEntry, data MessageIn) string {
	button.OnClicked(func(*ui.Button) {
		log.Println("Groups button listener function opened")
		output.SetText("")

		go func() {
			log.Println("Routine which is getting messages and opening group")
			for _, group := range data.GroupList {
				if group.GroupName == button.Text() {
					config.MessagesInGroup = group.Messages
					config.MembersInGroup = group.Members
					break
				}
			}
			for _, user := range config.MembersInGroup {
				config.UsersInGroup[user.ID] = user.Login
			}

			for _, message := range config.MessagesInGroup {
				var login string
				for id, name := range config.UsersInGroup {
					if message.MessageSenderID == id {
						login = name
					}
				}
				output.Append(login + ": " + message.Content + "\n")
			}
		}()

		var members []structure.User
		members = append(members, *NewUser(config.Login, "testPassword", config.Login, "test@test.com", true, "testUserIcon"))
		members = append(members, *NewUser(button.Text(), "testPassword", button.Text(), "test@test.com", true, "testUserIcon"))

		config.GroupName = button.Text()
		config.CurrentGroup = config.GroupID[config.GroupName]
		//формирование новой структуры на отправку на сервер,
		//заполнение текущего экземпляра требуемыми полями.
		user := NewUser(config.Login, "", config.Login, "test@test.com", true, "testUserIcon")
		group := NewGroup(user, config.GroupName, config.UserID, 1)
		msg := NewMessage(user, group, "", config.UserID, 1, "Text")
		message := NewMessageOut(user, &structure.User{}, group, msg, members, 1, 0, "GetGroup")
		_, err := conn.Write([]byte(JSONencode(*message)))
		if err != nil {
			log.Println(err)
		}
	})
	return config.GroupName
}

//ContactsAction is hanging listeners for contacts buttons
func ContactsAction(button *ui.Button, conn net.Conn, contacts *ui.Window, chat *ui.Window) {
	button.OnClicked(func(*ui.Button) {
		var members []structure.User
		members = append(members, *NewUser(config.Login, "", config.Login, "", true, ""))
		members = append(members, *NewUser(button.Text(), "", button.Text(), "", true, ""))
		config.GroupName = config.Login + button.Text()
		config.UserGroups = append(config.UserGroups, config.GroupName)
		user := NewUser(config.Login, "", config.Login, "test@test.com", true, "testUserIcon")
		group := NewGroup(user, config.GroupName, config.UserID, 1)
		message := NewMessageOut(user, &structure.User{}, group, &structure.Message{}, members, 1, 0, "CreateGroup")
		_, err := conn.Write([]byte(JSONencode(*message)))
		if err != nil {
			log.Println(err)
		}

	})
}
