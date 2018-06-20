package client

import (
	"fmt"
	"log"
	"net"

	"github.com/ProtonMail/ui"
	"go_messenger/desktop/structure"
)

var users []string
var login string

//Draw func which must configure connection and draw window
//with further hierarchy
func Draw() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Println(err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	err = ui.Main(func() {
		drawAuthWindow(conn)
	})
	if err != nil {
		log.Fatal(err)
	}
}

var userName string

func drawAuthWindow(conn net.Conn) {
	window := ui.NewWindow("Chat", 500, 500, false)
	loginInput := ui.NewEntry()

	passwordInput := ui.NewPasswordEntry()
	loginLabel := ui.NewLabel("Login")
	passwordLabel := ui.NewLabel("Password")
	signIn := ui.NewButton("Sign in!")
	signUp := ui.NewButton("Sign up!")
	topBox := ui.NewHorizontalBox()
	botBox := ui.NewHorizontalBox()
	middleBox := ui.NewHorizontalBox()
	fieldsBox := ui.NewVerticalBox()
	leftFieldBoxPadding := ui.NewVerticalBox()
	rightFieldBoxPadding := ui.NewVerticalBox()
	mainBox := ui.NewVerticalBox()
	fieldsBox.Append(loginLabel, false)
	fieldsBox.Append(loginInput, false)
	fieldsBox.Append(passwordLabel, false)
	fieldsBox.Append(passwordInput, false)
	fieldsBox.Append(signIn, false)
	fieldsBox.Append(signUp, false)
	middleBox.Append(leftFieldBoxPadding, true)
	middleBox.Append(fieldsBox, false)
	middleBox.Append(rightFieldBoxPadding, true)
	mainBox.Append(topBox, true)
	mainBox.Append(middleBox, true)
	mainBox.Append(botBox, true)
	window.SetChild(mainBox)
	window.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	window.Show()



	//обработчик кнопки входа, который отправляет запрос на получение всех юзеров в базе
	//для вывода и создание кнопок с ними
	signIn.OnClicked(func(*ui.Button) {
		//формирование новой структуры на отправку на сервер,
		//заполнение текущего экземпляра требуемыми полями.

		message := MessageOut{
			User:structure.User{
				Login:login,
				Password:passwordInput.Text(),
				Username:login,
				Email:"test@test.com",
				Status:true,
				UserIcon:"testUserIcon",
			},
			Contact:nil,
			Group:nil,
			Message:nil,
			Members:nil,
			RelationType:1,
			MessageLimit:1,
			Action:"GetUsers",
		}
		_, err := conn.Write([]byte(JSONencode(message)))
		if err != nil {
			log.Println(err)
		}
		login = loginInput.Text()
		window.Hide()
		drawChatWindow(conn)
		log.Println(users)
	})
	signUp.OnClicked(func(*ui.Button) {
		//формирование новой структуры на отправку на сервер,
		//заполнение текущего экземпляра требуемыми полями.

		message := MessageOut{
			User:structure.User{
				Login:login,
				Password:passwordInput.Text(),
				Username:login,
				Email:"test@test.com",
				Status:true,
				UserIcon:"testUserIcon",
			},
			Contact:nil,
			Group:nil,
			Message:nil,
			Members:nil,
			RelationType:1,
			MessageLimit:1,
			Action:"CreateUser",
		}
		_, err := conn.Write([]byte(JSONencode(message)))
		if err != nil {
			log.Println(err)
		}
		window.Hide()
		drawChatWindow(conn)
	})

	channel := make(chan bool)

	go func() {
		for {
			msg := JSONdecode(conn)
			for _, contacts := range msg.ContactList {
				users = append(users, contacts.Login)
			}
			log.Println(users, "READER")
			channel <- msg.Status
		}
	}()

}

var groupName string

func drawChatWindow(conn net.Conn) *ui.Window {
	fmt.Println(users, "chat window", login)
	window := ui.NewWindow(login, 500, 500, false)
	input := ui.NewEntry()
	input.SetText("message")
	send := ui.NewButton("Send")
	output := ui.NewMultilineNonWrappingEntry()
	output.SetReadOnly(true)
	mainBox := ui.NewHorizontalBox()
	usersBox := ui.NewVerticalBox()
	buttonUserSlice := make([]*ui.Button, 0)
	for _, user := range users {
		if user != "" && user != login {
			buttonWithUser := ui.NewButton(user)
			usersBox.Append(buttonWithUser, false)
			buttonUserSlice = append(buttonUserSlice, buttonWithUser)
		}
	}

	//sliceMembers := make([]string, 0)
	fmt.Println(buttonUserSlice)
	for i := 0; i < len(buttonUserSlice)-1; i++ {
		ListenerButton(i, buttonUserSlice[i], conn)
		output.SetText("")
	}
	//fmt.Println(buttonUserSlice, "slice buttons", buttonUserSlice[0].Text())
	messageBox := ui.NewVerticalBox()
	messageBox.Append(output, true)
	//messageBox.Append(user, false)
	messageBox.Append(input, false)
	messageBox.Append(send, false)
	mainBox.Append(usersBox, false)
	mainBox.Append(messageBox, true)
	go func() {
		for {
			msg := JSONdecode(conn)
			if msg.Content != "" {
				output.Append(msg.UserName + ": " + msg.Content + "\n")
			}
			fmt.Println(msg.Status)
		}
	}()
	send.OnClicked(func(*ui.Button) {
		//FIX SLICEMEMBER
		log.Println(groupName)
		output.Append(login + ": " + input.Text())

		//формирование новой структуры на отправку на сервер,
		//заполнение текущего экземпляра требуемыми полями.

		message := MessageOut{
			User:structure.User{
				Login:login,
				Password:"testPassword",
				Username:login,
				Email:"test@test.com",
				Status:true,
				UserIcon:"testUserIcon",
			},
			Contact:nil,
			Group:structure.Group{
				User:structure.User{
					Login:login,
					Password:"testPassword",
					Username:login,
					Email:"test@test.com",
					Status:true,
					UserIcon:"testUserIcon",
				},
				GroupType:structure.GroupType{
					Type:1,
				},
				GroupName:groupName,
				GroupOwnerID:123,
				GroupTypeID:1,
			},
			Message:structure.Message{
				User:structure.User{
					Login:login,
					Password:"testPassword",
					Username:login,
					Email:"test@test.com",
					Status:true,
					UserIcon:"testUserIcon",
				},
				Group:structure.Group{
					User:structure.User{
						Login:login,
						Password:"testPassword",
						Username:login,
						Email:"test@test.com",
						Status:true,
						UserIcon:"testUserIcon",
					},
					GroupType:structure.GroupType{
						Type:1,
					},
					GroupName:groupName,
					GroupOwnerID:123,
					GroupTypeID:1,
				},
			},
			Members:nil,
			RelationType:1,
			MessageLimit:1,
			Action:"SendMessageTo",
		}
		_, err := conn.Write([]byte(JSONencode(message)))
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

//func drawSignInErrorWindow(conn net.Conn) {
//	window := ui.NewWindow("Chat", 100, 100, false)
//	back := ui.NewButton("Back")
//	error := ui.NewLabel("Wrong login or password!")
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

//ListenerButton is hanging listeners for contact button
func ListenerButton(number int, button *ui.Button, conn net.Conn) string {
	button.OnClicked(func(*ui.Button) {
		sliceMembers := []string{login, button.Text()}
		groupName = login + button.Text()
		_, err := conn.Write([]byte(JSONencode(login, "", "",
			0, groupName, 1,
			login, sliceMembers, " ", " ", "",
			" ", " ", " ", true, " ", "CreateGroup")))
		if err != nil {
			log.Println(err)
		}
		fmt.Println(login, groupName, number, "graphic 131")
	})
	return groupName
}
