package dbservice
	
import (
	"fmt"

	"../../db"
	"../../models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func dbconnect() *gorm.DB{
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", db.DB_HOST, db.DB_PORT, db.DB_USER, db.DB_NAME, db.DB_PASSWORD, db.DB_SSLMODE)
	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
    return db
}

func CreateUser(login string, password string, username string,email string, status bool, usericon string ) bool{
	db := dbconnect()
	defer db.Close()
	user := models.User{Login: login, Password: password, Username: username, Email: email, Status: status, UserIcon: usericon}
	if db.NewRecord(user){
		db.Create(&user)
		return true
	}else{
		return false
	}
}

func CreateGroupType(groupType string) bool{
	db := dbconnect()
	defer db.Close()
	gtype := models.GroupType{Type: groupType}
	if db.NewRecord(gtype){
		db.Create(&gtype)
		return true
	}else{
		return false
	}
}
func CreateMessageType(messageType string) bool{
	db := dbconnect()
	defer db.Close()
	mtype := models.MessageContentType{Type: messageType}
	if db.NewRecord(mtype){
		db.Create(&mtype)
		return true
	}else{
		return false
	}
}

func CreateRelationType(relationType string) bool{
	db := dbconnect()
	defer db.Close()
	rtype := models.RelationType{Type: relationType}
	if db.NewRecord(rtype){
		db.Create(&rtype)
		return true
	}else{
		return false
	}
}

func CreateUserRelation(relatingUser string, relatedUser string, relationType uint) bool{
	db := dbconnect()
	defer db.Close()
	relatingU := models.User{}
	relatedU := models.User{}
	db.Where("username = ?", relatingUser).First(&relatingU)
	db.Where("username = ?", relatedUser).First(&relatedU)
	relation := models.UserRelation{RelatingUser:relatingU.ID, RelatedUser: relatedU.ID,RelationTypeID:relationType}
	if db.NewRecord(relation){
		db.Create(&relation)
		return true
	}else{
		return false
	}
}

func CreateGroup(groupName string, groupOwner string, groupType uint) bool{
	db := dbconnect()
	defer db.Close()
	owner := models.User{}
	db.Where("username = ?", groupOwner).First(&owner)
	group := models.Group{GroupName:groupName,GroupOwnerID: owner.ID,GroupTypeID:groupType}
	if db.NewRecord(group){
		db.Create(&group)
		return true
	}else{
		return false
	}
}

func AddMessage(content string, username string, groupName string, contentType uint) bool{
	db := dbconnect()
	defer db.Close()
	sender := models.User{}
	recipient := models.Group{}
	db.Where("username = ?", username).First(&sender)
	db.Where("group_name = ?", groupName).First(&recipient)
	message := models.Message{Content: content,MessageSenderID: sender.ID,MessageRecipientID: recipient.ID,MessageContentTypeID:contentType}
	if db.NewRecord(message){
		db.Create(&message)
		return true
	}else{
		return false
	}
}

func AddGroupMember(username string, groupName string, lastmessage string) bool{
	db := dbconnect()
	defer db.Close()
	user := models.User{}
	group := models.Group{}
	message := models.Message{}
	db.Where("username = ?", username).First(&user)
	db.Where("group_name = ?", groupName).First(&group)
	db.Where("content = ?", lastmessage).First(&message)
	member := models.GroupMember{UserID: user.ID,GroupID: group.ID,LastReadMessageID: message.ID}
	if db.NewRecord(member){
		db.Create(&member)
		return true
	}else{
		return false
	}

}
