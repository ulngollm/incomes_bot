package repository

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

var db *gorm.DB

func init() {
	connection, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{}); 
	if err != nil {
		fmt.Println(err)
	}

	db = connection
	err = db.AutoMigrate(&Transaction{})
	if err != nil {
		log.Fatal(err)
	}
}

