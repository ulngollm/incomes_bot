package repository

import (
	"fmt"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
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

