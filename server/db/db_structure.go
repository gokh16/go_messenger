package db

import (
	"log"

	"../models"
	"github.com/go-gormigrate/gormigrate"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateDatabase() {
	database, err := gorm.Open("postgres", "host=0.0.0.0 port=5434 user=postgres " +
		"dbname=golangDB password=golang sslmode=disable")
	if err != nil {
        log.Fatal(err)
    }
    if err = database.DB().Ping(); err != nil {
        log.Fatal(err)
    }

    database.LogMode(true)

	defer database.Close()

	 m := gormigrate.New(database, gormigrate.DefaultOptions, []*gormigrate.Migration{
	 	{
			ID: "0",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.Group_Type{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("group_types").Error
			},
		},
		{
			ID: "1",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.Message_Content_Type{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("message_content_types").Error
			},
		},
		{
			ID: "2",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.Relation_Type{}).Error
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
				database.CreateTable(&models.Message{})
				database.Model(&models.Message{}).AddForeignKey("message_sender_id", "users(id)", "RESTRICT", "RESTRICT")
				database.Model(&models.Message{}).AddForeignKey("message_recipient_id", "groups(id)", "RESTRICT", "RESTRICT")
				database.Model(&models.Message{}).AddForeignKey("message_content_type_id", "message_content_types(id)", "RESTRICT", "RESTRICT")
				return tx.AutoMigrate(&models.Message{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("messages").Error
			},
		},
		{
			ID: "5",
			Migrate: func(tx *gorm.DB) error {
				database.CreateTable(&models.Group{})
				database.Model(&models.Group{}).AddForeignKey("group_owner_id", "users(id)", "RESTRICT", "RESTRICT")
				database.Model(&models.Group{}).AddForeignKey("group_type_id", "group_types(id)", "RESTRICT", "RESTRICT")
				return tx.AutoMigrate(&models.Group{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("groups").Error
			},
		},
		{
			ID: "6",
			Migrate: func(tx *gorm.DB) error {
				database.CreateTable(&models.Group_Member{})
				database.Model(&models.Group_Member{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
				database.Model(&models.Group_Member{}).AddForeignKey("group_id", "groups(id)", "RESTRICT", "RESTRICT")
				database.Model(&models.Group_Member{}).AddForeignKey("last_read_message_id", "messages(id)", "RESTRICT", "RESTRICT")
				return tx.AutoMigrate(&models.Group_Member{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("group_members").Error
			},
		},
		{
			ID: "7",
			Migrate: func(tx *gorm.DB) error {
				database.CreateTable(&models.User_Relation{})
				database.Model(&models.User_Relation{}).AddForeignKey("relation_type_id", "relation_types(id)", "RESTRICT", "RESTRICT")				

				return tx.AutoMigrate(&models.User_Relation{}).Error
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

















