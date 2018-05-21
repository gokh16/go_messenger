package dbservice
	
import (
	"../../models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func dbconnect() *gorm.DB{
	database, err := gorm.Open("postgres", "host=0.0.0.0 port=5434 user=postgres " +
		"dbname=golangDB password=golang sslmode=disable")
	if err != nil {
        panic(err)
    }
    if err = database.DB().Ping(); err != nil {
        panic(err)
    }
    return database
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

func CreateGroupType(group_type string) bool{
	db := dbconnect()
	defer db.Close()
	gtype := models.GroupType{Type: group_type}
	if db.NewRecord(gtype){
		db.Create(&gtype)
		return true
	}else{
		return false
	}
}
func CreateMessageType(message_type string) bool{
	db := dbconnect()
	defer db.Close()
	mtype := models.MessageContentType{Type: message_type}
	if db.NewRecord(mtype){
		db.Create(&mtype)
		return true
	}else{
		return false
	}
}

func CreateRelationType(relation_type string) bool{
	db := dbconnect()
	defer db.Close()
	rtype := models.RelationType{Type: relation_type}
	if db.NewRecord(rtype){
		db.Create(&rtype)
		return true
	}else{
		return false
	}
}

func CreateUserRelation(relating_user string, related_user string, relation_type uint) bool{
	db := dbconnect()
	defer db.Close()
	relating_u := models.User{}
	related_u := models.User{}
	db.Where("username = ?", relating_user).First(&relating_u)
	db.Where("username = ?", related_user).First(&related_u)
	relation := models.UserRelation{RelatingUser:relating_u.ID, RelatedUser: related_u.ID,RelationTypeID:relation_type}
	if db.NewRecord(relation){
		db.Create(&relation)
		return true
	}else{
		return false
	}
}

func CreateGroup(group_name string, group_owner string, group_type uint) bool{
	db := dbconnect()
	defer db.Close()
	owner := models.User{}
	db.Where("username = ?", group_owner).First(&owner)
	group := models.Group{GroupName:group_name,GroupOwnerID: owner.ID,GroupTypeID:group_type}
	if db.NewRecord(group){
		db.Create(&group)
		return true
	}else{
		return false
	}
}

func AddMessage(content string, username string, group_name string, content_type uint) bool{
	db := dbconnect()
	defer db.Close()
	sender := models.User{}
	recipient := models.Group{}
	db.Where("username = ?", username).First(&sender)
	db.Where("group_name = ?", group_name).First(&recipient)
	message := models.Message{Content: content,MessageSenderID: sender.ID,MessageRecipientID: recipient.ID,MessageContentTypeID:content_type}
	if db.NewRecord(message){
		db.Create(&message)
		return true
	}else{
		return false
	}
}

func AddGroupMember(username string, group_name string, lastmessage string) bool{
	db := dbconnect()
	defer db.Close()
	user := models.User{}
	group := models.Group{}
	message := models.Message{}
	db.Where("username = ?", username).First(&user)
	db.Where("group_name = ?", group_name).First(&group)
	db.Where("content = ?", lastmessage).First(&message)
	member := models.GroupMember{UserID: user.ID,GroupID: group.ID,LastReadMessageID: message.ID}
	if db.NewRecord(member){
		db.Create(&member)
		return true
	}else{
		return false
	}

}