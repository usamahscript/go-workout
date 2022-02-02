//Purpose of this file is to return a variabel called DB which will allow other files ot interact with the DB

package config

import (
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialect/mysql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Connect() {
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbname := os.Getenv("NAME")
	dbpassword := os.Getenv("PASSWORD")
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbname, dbpassword, dbPort)
	db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully Connected to the Database")
	}
	// db
}

func GetDB() *gorm.DB {
	return db
}
