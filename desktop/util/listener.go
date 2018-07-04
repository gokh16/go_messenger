package util

import (
	"log"
	"net"

	"go_messenger/desktop/structure"

	"go_messenger/desktop/config"

	"github.com/ProtonMail/ui"
)

//ButtonActions is hanging listeners for group buttons
func ButtonActions(button *ui.Button, conn net.Conn, output *ui.MultilineEntry) string {
	button.OnClicked(func(*ui.Button) {
		output.SetText("")

		go func() {
			users := make(map[uint]string)
			members := make([]structure.User, 0)
			for {
				msg := JSONdecode(conn)
				for _, group := range msg.GroupList {
					if group.GroupName == button.Text() {
						config.MessagesInGroup = group.Messages
						members = group.Members
						break
					}
				}
				for _, user := range members {
					users[user.ID] = user.Login
				}
				break
			}
			for _, message := range config.MessagesInGroup {
				var login string
				for id, name := range users {
					if message.MessageSenderID == id {
						login = name
					}
				}
				output.Append(login + ": " + message.Content + "\n")
			}
			config.MarkForRead <- "chat"
		}()

		var members []structure.User
		members = append(members, structure.User{
			Login:    config.Login,
			Password: "testPassword",
			Username: config.Login,
			Email:    "test@test.com",
			Status:   true,
			UserIcon: "testUserIcon",
		})
		members = append(members, structure.User{
			Login:    button.Text(),
			Password: "testPassword",
			Username: button.Text(),
			Email:    "test@test.com",
			Status:   true,
			UserIcon: "testUserIcon",
		})

		config.GroupName = button.Text()
		//формирование новой структуры на отправку на сервер,
		//заполнение текущего экземпляра требуемыми полями.
		user := NewUser(config.Login, "", config.Login, "test@test.com", true, "testUserIcon")
		group := NewGroup(user, "private", config.GroupName, config.UserID, 1)
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
		config.MarkForRead <- "contacts"

		var members []structure.User
		members = append(members, *NewUser(config.Login, "", config.Login, "", true, ""))
		members = append(members, *NewUser(button.Text(), "", button.Text(), "", true, ""))
		config.GroupName = config.Login + button.Text()
		config.UserGroups = append(config.UserGroups, config.GroupName)
		user := NewUser(config.Login, "", config.Login, "test@test.com", true, "testUserIcon")
		group := NewGroup(user, "private", config.GroupName, config.UserID, 1)
		message := NewMessageOut(user, &structure.User{}, group, &structure.Message{}, members, 1, 0, "CreateGroup")
		_, err := conn.Write([]byte(JSONencode(*message)))
		if err != nil {
			log.Println(err)
		}
		go func() {
			status := <-config.MarkForRead
			for {
				if status == "contacts" {
					msg := JSONdecode(conn)
					if msg.Status{
						config.MarkForRedrawChatWindow <- "groups are accepted"
					}
				}

			}
		}()

	})
}
