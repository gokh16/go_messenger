package orm

import (
	"log"

	"github.com/go-gormigrate/gormigrate"
	"github.com/jinzhu/gorm"
	//ignoring init from package below
	"go_messenger/server/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//CreateDatabase method which establish connections to database
func CreateDatabase() {
	db, err := gorm.Open("postgres", "host=0.0.0.0 port=5432 user=postgres "+
		"dbname=golangDB password=root sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.DB().Ping(); err != nil {
		log.Fatal(err)
	}

	db.LogMode(true)

	defer func() {
		err := db.Close()
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
			//todo missing return, ask to Max
		},
		{
			ID: "1",
			Migrate: func(tx *gorm.DB) error {
				//todo resolve this problem, ask to Max
				//return tx.AutoMigrate(&models.MessageContentType{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("message_content_types").Error
			},
		},
		{
			ID: "2",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.RelationType{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("relation_types").Error
			},
		},
		{
			ID: "3",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.User{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("users").Error
			},
		},
		{
			ID: "4",
			Migrate: func(tx *gorm.DB) error {
				db.CreateTable(&models.Message{})
				db.Model(&models.Message{}).AddForeignKey("message_sender_id", "users(id)", "RESTRICT", "RESTRICT")
				db.Model(&models.Message{}).AddForeignKey("message_recipient_id", "groups(id)", "RESTRICT", "RESTRICT")
				db.Model(&models.Message{}).AddForeignKey("message_content_type_id", "message_content_types(id)", "RESTRICT", "RESTRICT")
				return tx.AutoMigrate(&models.Message{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("messages").Error
			},
		},
		{
			ID: "5",
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
			ID: "6",
			Migrate: func(tx *gorm.DB) error {
				db.CreateTable(&models.GroupMember{})
				db.Model(&models.GroupMember{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
				db.Model(&models.GroupMember{}).AddForeignKey("group_id", "groups(id)", "RESTRICT", "RESTRICT")
				db.Model(&models.GroupMember{}).AddForeignKey("last_read_message_id", "messages(id)", "RESTRICT", "RESTRICT")
				return tx.AutoMigrate(&models.GroupMember{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("group_members").Error
			},
		},
		{
			ID: "7",
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
