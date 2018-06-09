package dbservice

import (
	"fmt"
	"go_messenger/server/db"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var dbConn *gorm.DB

// OpenConnDB opens a connection eith DB it returns a *DB object for closing the connection in main
func OpenConnDB() *gorm.DB {
	var err error
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", db.DB_HOST, db.DB_PORT, db.DB_USER, db.DB_NAME, db.DB_PASSWORD, db.DB_SSLMODE)
	dbConn, err = gorm.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("gorm Open connection error: ", err)
	}
	fmt.Println("db open ok")
	return dbConn
}
