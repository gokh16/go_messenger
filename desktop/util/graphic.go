package util

import (
	"fmt"
	"log"
	"net"

	"go_messenger/desktop/structure"

	"github.com/ProtonMail/ui"
	"go_messenger/desktop/config"
)

//ListenerButton is hanging listeners for contact button
func ListenerButton(number int, button *ui.Button, conn net.Conn) string {
	button.OnClicked(func(*ui.Button) {

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

		config.GroupName = config.Login + button.Text()

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
				GroupName:    config.GroupName,
				GroupOwnerID: 123,
				GroupTypeID:  1,
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
					GroupName:    config.GroupName,
					GroupOwnerID: 123,
					GroupTypeID:  1,
				},
			},
			Members:      members,
			RelationType: 1,
			MessageLimit: 1,
			Action:       "SendMessageTo",
		}
		_, err := conn.Write([]byte(JSONencode(message)))
		if err != nil {
			log.Println(err)
		}
		fmt.Println(config.Login, config.GroupName, number, "graphic 131")
	})
	return config.GroupName
}

//func drawSignInErrorWindow(conn net.Conn) {
//	window := ui.NewWindow("Chat", 100, 100, false)
//	back := ui.NewButton("Back")
//	error := ui.NewLabel("Wrong Login or password!")
//	box := ui.NewVerticalBox()
//	box.Append(back, false)
//	box.Append(error, false)
//	window.SetChild(box)
//	back.OnClicked(func(*ui.Button) {
//		drawAuthWindow(conn)
//		window.Hide()
//	})
//	window.Show()
//}

//func GetUser(conn net.Conn) []string {
//	conn.Write([]byte(JSONencode("", "", "",
//		0, " ", 1,
//		" ", nil, " ", "", "",
//		" ", " ", " ", true, " ", "GetUsers")))
//	time.Sleep(2 * time.Second)
//	msg := JSONdecode(conn)
//	return msg.GroupMember
//}
