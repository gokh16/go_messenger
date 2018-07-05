package windows

import (
	"go_messenger/desktop/config"
	"go_messenger/desktop/util"
	"log"
	"net"
)

var InputData = make(chan util.MessageIn, 1024)

//Reader method is listening connection and routes data to the next step
func Reader(conn net.Conn) {
	for {
		msg := util.JSONdecode(conn)
		log.Println(msg.Action, "reader")
		switch msg.Action {
		case "LoginUser":
			InputData <- msg
			for _, contacts := range msg.GroupList {
				config.UserGroups = append(config.UserGroups, contacts.GroupName)
				config.GroupID[contacts.GroupName] = contacts.ID
			}
			log.Println("here")
			config.UserID = msg.User.ID
		case "SendMessageTo":
			InputData <- msg

		default:
			log.Println("default")
		}
	}
}
