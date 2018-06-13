package db

import (
	"fmt"
	"log"

	"go_messenger/server/models"

	"github.com/go-gormigrate/gormigrate"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//CreateDatabase creates tables and relations in DB
func CreateDatabase() {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", DB_HOST, DB_PORT, DB_USER, DB_NAME, DB_PASSWORD, DB_SSLMODE)
	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "0",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.GroupType{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("group_types").Error
			},
		},
		{
			ID: "1",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.RelationType{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("relation_types").Error
			},
		},
		{
			ID: "2",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.User{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("users").Error
			},
		},
		{
			ID: "3",
			Migrate: func(tx *gorm.DB) error {
				db.CreateTable(&models.Message{})
				db.Model(&models.Message{}).AddForeignKey("message_sender_id", "users(id)", "RESTRICT", "RESTRICT")
				db.Model(&models.Message{}).AddForeignKey("message_recipient_id", "groups(id)", "RESTRICT", "RESTRICT")
				return tx.AutoMigrate(&models.Message{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("messages").Error
			},
		},
		{
			ID: "4",
			Migrate: func(tx *gorm.DB) error {
				db.CreateTable(&models.Group{})
				db.Model(&models.Group{}).AddForeignKey("group_owner_id", "users(id)", "RESTRICT", "RESTRICT")
				db.Model(&models.Group{}).AddForeignKey("group_type_id", "group_types(id)", "RESTRICT", "RESTRICT")
				return tx.AutoMigrate(&models.Group{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("groups").Error
			},
		},
		{
			ID: "5",
			Migrate: func(tx *gorm.DB) error {
				db.CreateTable(&models.GroupMember{})
				db.Model(&models.GroupMember{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
				db.Model(&models.GroupMember{}).AddForeignKey("group_id", "groups(id)", "RESTRICT", "RESTRICT")
				return tx.AutoMigrate(&models.GroupMember{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("group_members").Error
			},
		},
		{
			ID: "6",
			Migrate: func(tx *gorm.DB) error {
				db.CreateTable(&models.UserRelation{})
				db.Model(&models.UserRelation{}).AddForeignKey("relation_type_id", "relation_types(id)", "RESTRICT", "RESTRICT")

				return tx.AutoMigrate(&models.UserRelation{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("user_relations").Error
			},
		},
	})
	err = m.Migrate()
	if err == nil {
		log.Printf("Migration did run successfully")
	} else {
		log.Printf("Could not migrate: %v", err)
	}
}

func InitDatabase() {
	//for testing
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", DB_HOST, DB_PORT, DB_USER, DB_NAME, DB_PASSWORD, DB_SSLMODE)
	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	group := models.Group{GroupName: "User1User3", GroupOwnerID: 1, GroupTypeID: 1}
	db.Create(&group)
	groupmember1 := models.GroupMember{GroupID: 2, UserID: 1, LastReadMessageID: 1}
	groupmember2 := models.GroupMember{GroupID: 2, UserID: 3, LastReadMessageID: 1}
	db.Create(&groupmember1)
	db.Create(&groupmember2)
	/*type1 := models.GroupType{Type: "Private"}
	db.Create(&type1)
	user1 := models.User{Login: "User1", Password: "", Username: "User1", Status: false, UserIcon: ""}
	user2 := models.User{Login: "User2", Password: "", Username: "User2", Status: false, UserIcon: ""}
	group := models.Group{GroupName: "User1User2", GroupOwnerID: 1, GroupTypeID: 1}
	db.Create(&user1)
	db.Create(&user2)
	db.Create(&group)
	msg := models.Message{Content: "Hello", MessageContentType: "Text", MessageRecipientID: 1, MessageSenderID: 2}
	db.Create(&msg)
	groupmember1 := models.GroupMember{GroupID: 1, UserID: 1, LastReadMessageID: 1}
	groupmember2 := models.GroupMember{GroupID: 1, UserID: 2, LastReadMessageID: 1}
	db.Create(&groupmember1)
	db.Create(&groupmember2)*/
}
