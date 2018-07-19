package config

import (
	"go_messenger/desktop/structure"
)

var UserGroups []string
var Login string
var Password string
var GroupName string
var UserID uint
var MessagesInGroup []structure.Message
var GroupID = make(map[string]uint)
var CurrentGroup uint
var UsersInGroup = make(map[uint]string)
var MembersInGroup []structure.User
