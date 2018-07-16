package windows

import (
	"go_messenger/desktop/config"
	"go_messenger/desktop/structure"
	"go_messenger/desktop/util"
	"log"
	"net"

	"github.com/ProtonMail/ui"
)

//DrawAuthWindow is a func which draw window by GTK's help
func DrawAuthWindow(conn net.Conn) {
	log.Println("Opened DrawAuthWindow")
	window := ui.NewWindow("Humble", 500, 500, false)
	loginInput := ui.NewEntry()
	passwordInput := ui.NewPasswordEntry()
	loginLabel := ui.NewLabel("Login")
	passwordLabel := ui.NewLabel("Password")
	signIn := ui.NewButton("Sign in!") //asd
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
		log.Println("Button Beginning clicked")
		config.Login = loginInput.Text()
		//формирование новой структуры на отправку на сервер,
		//заполнение текущего экземпляра требуемыми полями.
		user := util.NewUser(config.Login, passwordInput.Text(), config.Login, "test@test.com", true, "testUserIcon")
		message := util.NewMessageOut(user, &structure.User{}, &structure.Group{}, &structure.Message{}, nil, 1, 0, "LoginUser")

		_, err := conn.Write([]byte(util.JSONencode(*message)))
		if err != nil {
			log.Println(err)
		}
		if passwordInput.Text() != "" || loginInput.Text() != "" {
			window.Hide()
			DrawChatWindow(conn)
		} else {
			window.Hide()
			DrawErrorWindow("Enter the password!", conn)
		}
		return
	})
	signUp.OnClicked(func(*ui.Button) {
		log.Println("Button SignUp clicked")
		DrawRegistrationWindow(conn)
		window.Hide()
		return
	})

}
