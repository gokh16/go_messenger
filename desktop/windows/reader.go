package windows

import (
	"go_messenger/desktop/config"
	"go_messenger/desktop/structure"
	"go_messenger/desktop/util"
	"net"
)

//var InputData = make(chan util.MessageIn)
var Contacts = make(chan []structure.User)
var Send = make(chan util.MessageIn)
var Beginning = make(chan util.MessageIn)
var SignUp = make(chan util.MessageIn)
var Groups = make(chan util.MessageIn)

//Reader method is listening connection and routes data to the next step
func Reader(conn net.Conn) {
	for {
		msg := util.JSONdecode(conn)
		switch msg.Action {
		case "LoginUser":
			Beginning <- msg
			for _, contacts := range msg.GroupList {
				config.UserGroups = append(config.UserGroups, contacts.GroupName)
				config.GroupID[contacts.GroupName] = contacts.ID
			}
			config.UserID = msg.User.ID
		case "CreateUser":
			SignUp <- msg
		case "GetUser":
			Contacts <- msg.ContactList
		case "SendMessageTo":
			Send <- msg
		case "GetContactList":
			Contacts <- msg.ContactList
		case "GetGroupList":
			Beginning <- msg
		}
	}
}
