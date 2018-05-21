package main

import (
	"./db/migrations"
	"./db"
)
func init(){
	migrations.CreateDatabase();
}

func main(){
	dbservice.CreateUser("0969769486","123456","_lunarmax","https://pp.userapi.com/c847017/v847017389/1ddb4/OidCN-HrCx4.jpg")
	dbservice.CreateGroupType("private_message")
	dbservice.CreateGroup("Group1", "_lunarmax", 1)
	dbservice.CreateMessageType("text")
	dbservice.AddMessage("Hello","_lunarmax","Group1",1)
}