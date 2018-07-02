package util

import (
	"log"
	"net"

	"go_messenger/desktop/structure"

	"go_messenger/desktop/config"

	"github.com/ProtonMail/ui"
)

//ButtonActions is hanging listeners for contact button
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
			config.MarkForRead <- true
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

		message := MessageOut{
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
				//GroupOwnerID: 123,
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
			},
			Members:      members,
			RelationType: 1,
			MessageLimit: 10,
			Action:       "GetGroup",
		}
		_, err := conn.Write([]byte(JSONencode(message)))
		if err != nil {
			log.Println(err)
		}
	})
	return config.GroupName
}
