package windows

import (
	"net"

	"go_messenger/desktop/structure"
	"go_messenger/desktop/util"
	"log"

	"regexp"

	"github.com/ProtonMail/ui"
)

func DrawRegistrationWindow(conn net.Conn) {
	window := ui.NewWindow("Humble", 500, 500, false)
	loginInput := ui.NewEntry()
	passwordInput := ui.NewPasswordEntry()
	emailInput := ui.NewEntry()
	usernameInput := ui.NewEntry()
	loginLabel := ui.NewLabel("Login")
	passwordLabel := ui.NewLabel("Password")
	emailLabel := ui.NewLabel("Email")
	usernameLabel := ui.NewLabel("Username")
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
	fieldsBox.Append(emailLabel, false)
	fieldsBox.Append(emailInput, false)
	fieldsBox.Append(usernameLabel, false)
	fieldsBox.Append(usernameInput, false)
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

	signUp.OnClicked(func(*ui.Button) {
		//формирование новой структуры на отправку на сервер,
		//заполнение текущего экземпляра требуемыми полями.
		user := util.NewUser(loginInput.Text(), passwordInput.Text(), usernameInput.Text(), emailInput.Text(), true, "testUserIcon")
		message := util.NewMessageOut(user, &structure.User{}, &structure.Group{}, &structure.Message{}, nil, 1, 0, "CreateUser")
		_, err := conn.Write([]byte(util.JSONencode(*message)))
		if err != nil {
			log.Println(err)
		}
		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if !re.MatchString(emailInput.Text()) {
			window.Hide()
			DrawErrorWindow("Invalid email!", conn)
			return
		}
		if passwordInput.Text() != "" || loginInput.Text() != "" || emailInput.Text() != "" {
			window.Hide()
			DrawAuthWindow(conn)
		} else {
			window.Hide()
			DrawErrorWindow("You need to fill fields first!", conn)
		}
	})
	window.Show()
}
