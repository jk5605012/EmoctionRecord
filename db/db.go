package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// test
var DB *gorm.DB

// func Init() {
// 	DB, err := gorm.Open()
// }
