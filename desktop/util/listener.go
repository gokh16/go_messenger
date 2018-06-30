package util

import (
	"fmt"
	"log"
	"net"

	"go_messenger/desktop/structure"

	"github.com/ProtonMail/ui"
	"go_messenger/desktop/config"
)

//ButtonActions is hanging listeners for contact button
func ButtonActions(button *ui.Button, conn net.Conn, output *ui.MultilineEntry) string {
	button.OnClicked(func(*ui.Button) {
		output.SetText("")
		output.SetText("Now you can texting with:" + button.Text() + "\n") //todo fix group messaging

		go func() {
			for {
				msg := JSONdecode(conn)
				for _, group := range msg.GroupList {
					if msg.Message.Content != "" && group.GroupName == button.Text() {
						config.MessagesInGroup = group.Messages
						break
					}
				}
				break
				fmt.Println(msg.Status)
			}
			for _, message := range config.MessagesInGroup {
				output.Append(message.User.Login + ": " + message.Content + "\n")
			}
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
			MessageLimit: 5,
			Action:       "GetGroup",
		}
		_, err := conn.Write([]byte(JSONencode(message)))
		if err != nil {
			log.Println(err)
		}
	})
	return config.GroupName
}