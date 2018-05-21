package main

import (
	"./db"
	"./db/dbservice"
)
func init(){
	db.CreateDatabase();
}

func main(){

	dbservice.CreateMessageType("text")
	dbservice.CreateRelationType("friends")
	dbservice.CreateGroupType("private_message")

	dbservice.CreateUser("0969769486","123456","_lunarmax","just_my_mail@coolsite.net",false,"https://pp.userapi.com/c847017/v847017389/1ddb4/OidCN-HrCx4.jpg")
	dbservice.CreateUser("0969769486","123456","_lunarlexy","just_my_mail@coolsite.net",false,"https://pp.userapi.com/c847017/v847017389/1ddb4/OidCN-HrCx4.jpg")
	
	dbservice.CreateGroup("Group1", "_lunarmax", 1)
	
	dbservice.AddMessage("Hello lexy","_lunarmax","Group1",1)
	
	dbservice.AddGroupMember("_lunarlexy","Group1","Hello lexy")
	
	dbservice.AddMessage("I am your friend","_lunarmax","Group1",1)
	
	dbservice.CreateUserRelation("_lunarmax","_lunarlexy",1)
}