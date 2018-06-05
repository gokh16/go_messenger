package service

import (
	"go_messenger/server/db/dbservice"
	"go_messenger/server/service/interfaces"
)

//AddGroupMember add new members in spesific Group.
func AddGroupMember(userName, groupName, lastMessage string, groupMember []string) {
	var gmi interfaces.GMI = dbservice.GroupMember{}
	for _, user := range groupMember {
		gmi.AddGroupMember(user, groupName, lastMessage)
	}
}
