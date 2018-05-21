package main

import (
	"./db"
	"./db/dbservice"
	"fmt"
)

func init(){
	db.CreateDatabase();
}

func main(){

	/*dbservice.CreateMessageType("text")
	dbservice.CreateRelationType("friends")
	dbservice.CreateGroupType("private_message")

	dbservice.CreateUser("0969769486","123456","_lunarmax","just_my_mail@coolsite.net",false,"https://pp.userapi.com/c847017/v847017389/1ddb4/OidCN-HrCx4.jpg")
	dbservice.CreateUser("0969769486","123456","_lunarlexy","just_my_mail@coolsite.net",false,"https://pp.userapi.com/c847017/v847017389/1ddb4/OidCN-HrCx4.jpg")
	
	dbservice.CreateGroup("Group1", "_lunarmax", 1)
	
	dbservice.AddMessage("Hello lexy","_lunarmax","Group1",1)
	
	dbservice.AddGroupMember("_lunarlexy","Group1","Hello lexy")
	
	dbservice.AddMessage("I am your friend","_lunarmax","Group1",1)
	dbservice.AddMessage("Test1","_lunarmax","Group1",1)
	dbservice.AddMessage("Test2","_lunarmax","Group1",1)
	dbservice.AddMessage("Test3","_lunarmax","Group1",1)
	dbservice.AddMessage("Test4","_lunarmax","Group1",1)
	dbservice.AddMessage("Test5","_lunarmax","Group1",1)*/

	u := dbservice.GetContactList("_lunarmax")
	for i,_:= range u{
		fmt.Println(u[i].Username)
	}
}