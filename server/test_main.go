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

	
	dbservice.CreateMessageType("text")
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
	dbservice.AddMessage("Test5","_lunarmax","Group1",1)
	
	
	dbservice.CreateUserRelation("_lunarmax","_lunarlexy",1)
	dbservice.CreateUserRelation("_lunarlexy","_lunarmax",1)
	
	dbservice.CreateUser("0953644890","qwerty123","Lyxid","lyxid@gmail.com",false,"htttps:")
	dbservice.CreateUserRelation("_lunarlexy","Lyxid",1)
	
	fmt.Println("GetMessages(5,'Group1')")
	u := dbservice.GetMessages(5,"Group1")
	for i,_:= range u{
		fmt.Println(u[i].Content)
	}
	fmt.Println()
	fmt.Println("GetGroupList('_lunarlexy')")
	u1 := dbservice.GetGroupList("_lunarlexy")
	for i,_:= range u1{
		fmt.Println(u1[i].GroupName)
	}
	fmt.Println()
	fmt.Println("GetGroupUserList('Group1')")
	u2 := dbservice.GetGroupUserList("Group1")
	for i,_:= range u2{
		fmt.Println(u2[i].Username)
	}
	fmt.Println()
	fmt.Println("FindUser('_lunarmax')")
	u3 := dbservice.FindUser("_lunarmax")
	fmt.Println(u3.Username)
	
	fmt.Println()
	fmt.Println("GetContactList('_lunarlexy')")
	u4:= dbservice.GetContactList("_lunarlexy")
	for i,_:= range u4{
		fmt.Println(u4[i].Username)
	}
}
