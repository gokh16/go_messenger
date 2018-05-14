package main

import (
	"log"
	"fmt"

	"github.com/go-gormigrate/gormigrate"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=0.0.0.0 port=5434 user=postgres " +
		"dbname=golangDB password=golang sslmode=disable")
	if err != nil {
        log.Fatal(err)
    }
    if err = db.DB().Ping(); err != nil {
        log.Fatal(err)
    }

    db.LogMode(true)

	defer db.Close()

	 m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
	 	{
			ID: "0",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Group_Type{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("group_types").Error
			},
		},
		{
			ID: "1",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Message_Content_Type{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("message_content_types").Error
			},
		},
		{
			ID: "2",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Relation_Type{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("relation_types").Error
			},
		},
		{
			ID: "3",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&User{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("users").Error
			},
		},
		{
			ID: "4",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Message{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("messages").Error
			},
		},
		{
			ID: "5",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Group{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("groups").Error
			},
		},
		{
			ID: "6",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Group_Member{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("group_members").Error
			},
		},
		{
			ID: "7",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&User_Relation{}).Error
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


type User struct {
	gorm.Model

	Login string
	Password string
	Username string
	UserIcon string
}

type Message struct {
	gorm.Model

	
	Content string
	Message_sender int  `sql:"type:int REFERENCES users(id)"`
	Message_recepient int  `sql:"type:int REFERENCES groups(id)"`
	Message_content_type int `sql:"type:int REFERENCES message_content_types(id)"`
}

type Group struct {
	gorm.Model

	Group_owner int `sql:"type:int REFERENCES users(id)"`
	Group_type int `sql:"type:int REFERENCES groups(id)"`
}

type Group_Member struct {
	gorm.Model

	User_id int `sql:"type:int REFERENCES users(id)"`
	Group_id int `sql:"type:int REFERENCES groups(id)"`
	Last_read_message_id int `sql:"type:int REFERENCES messages(id)"`
}

type Message_Content_Type struct {
	gorm.Model

	Type string
}

type Group_Type struct {
	gorm.Model

	Type string
}

type User_Relation struct {
	gorm.Model

	Relating_user int `sql:"type:int REFERENCES users(id)"`
	Related_user int `sql:"type:int REFERENCES users(id)"`
	Relation_type int `sql:"type:int REFERENCES relation_types(id)"`
}

type Relation_Type struct {
	gorm.Model

	Type string
}
