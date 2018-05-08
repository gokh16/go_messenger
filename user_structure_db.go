package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
)

const (
	host = "0.0.0.0"
	port = 5434
	user = "golang"
	password = "golang"
	dbname = "golang"
)

func main() {
	db, err := gorm.Open("postgres", "host=0.0.0.0 port=5434 user=golang " +
		"dbname=golang password=golang sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.CreateTable(&User{})
	fmt.Println("Successfully connected!")
}

type User struct {
	gorm.Model

	Username string
	Password string
	Name string
	UserIcon string
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %v>", u.ID, u.Name, u.Username)
}
