package dbservice

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var conn *gorm.DB
var err error

// OpenConnDB opens a connection eith DB it returns a *DB object for closing the connection in main
func OpenConnDB() *gorm.DB {
	conn, err = gorm.Open("postgres", "user=postgres password=1111 dbname=postgres sslmode=disable")
	if err != nil {
		fmt.Println("gorm Open connection error: ", err)
	}
	return conn
}
