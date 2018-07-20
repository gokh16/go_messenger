package windows

import (
	"go_messenger/desktop/config"
	"go_messenger/desktop/structure"
	"go_messenger/desktop/util"
	"log"
	"net"
)

var Contacts = make(chan []structure.User)
var Send = make(chan util.MessageIn)
var Beginning = make(chan util.MessageIn)
var SignUp = make(chan util.MessageIn)

//Reader method is listening connection and routes data to the next step
func Reader(conn net.Conn) {
	for {
		msg := util.JSONdecode(conn)
		switch msg.Action {
		case "LoginUser":
			Beginning <- msg
			config.UserGroups = nil
			for _, contacts := range msg.GroupList {
				config.UserGroups = append(config.UserGroups, contacts.GroupName)
				log.Println(contacts.ID)
				config.GroupID[contacts.GroupName] = contacts.ID
			}
			config.UserID = msg.User.ID
			//fmt.Println(config.GroupID)
		case "CreateUser":
			SignUp <- msg
		case "GetUser":
			Contacts <- msg.ContactList
		case "SendMessageTo":
			Send <- msg
			log.Println(msg.Message.MessageRecipientID)
		case "GetContactList":
			Contacts <- msg.ContactList
		case "GetGroupList":
			Beginning <- msg
		}
	}
}
