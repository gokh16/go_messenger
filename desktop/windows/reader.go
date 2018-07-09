package windows

import (
	"go_messenger/desktop/config"
	"go_messenger/desktop/util"
	"log"
	"net"
)

var InputData = make(chan util.MessageIn)

//Reader method is listening connection and routes data to the next step
func Reader(conn net.Conn) {
	for {
		msg := util.JSONdecode(conn)
		InputData <- msg
		switch msg.Action {
		case "LoginUser":
			for _, contacts := range msg.GroupList {
				config.UserGroups = append(config.UserGroups, contacts.GroupName)
				config.GroupID[contacts.GroupName] = contacts.ID
			}
			config.UserID = msg.User.ID
		default:
			log.Println("default")
		}
	}
}
