package db

import (
	"fmt"
	"go_messenger/server/models"
	"log"

	"github.com/go-gormigrate/gormigrate"
	"github.com/jinzhu/gorm"
	//ignoring init from package below
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//CreateDatabase creates tables and relations in DB
func CreateDatabase() {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", HostDB, PortDB, UserDB, NameDB, PasswordDB, SSLModeDB)
	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		log.Println(err)
	}
	defer func() {
		err = db.Close()
		if err != nil {
			log.Println(err)
		}
	}()

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
	//initGroupType(db)
}

//func initGroupType(db *gorm.DB) {
//	private := models.GroupType{Type: "private"}
//	private.ID = 1
//	public := models.GroupType{Type: "public"}
//	public.ID = 2
//	if !db.NewRecord(&private) {
//		db.Create(&private)
//	}
//	if !db.NewRecord(&public) {
//		db.Create(&public)
//	}
//}

