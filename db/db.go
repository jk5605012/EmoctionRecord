package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// test
var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open("postgres", "postgres://postgres:mysecretpassword@localhost:5432/test?sslmode=disable")
	if err != nil {
		log.Fatal(err)
		return
	}
}
