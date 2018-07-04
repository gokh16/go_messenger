package dbservice

import (
	"fmt"
	"go_messenger/server/db"

	"github.com/jinzhu/gorm"
	//ignoring init from package below
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var dbConn *gorm.DB

// OpenConnDB opens a connection either DB it returns a *DB object for closing the connection in main
func OpenConnDB() *gorm.DB {
	var err error
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", db.HostDB, db.PortDB, db.UserDB, db.NameDB, db.PasswordDB, db.SSLModeDB)
	dbConn, err = gorm.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("gorm Open connection error: ", err)
	}
	fmt.Println("db open ok")
	return dbConn
}
