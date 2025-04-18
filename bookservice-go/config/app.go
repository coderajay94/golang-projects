package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// initializing the database connection in gorm v2 is different from v1
// gorm v1 is with github.com/jinzhu/gorm
// gorm v2 is with gorm.io/gorm
func ConnectMySQLDB() {
	dsn := "admin:password@tcp(127.0.0.1:3306)/simplerest?charset=utf8mb4&parseTime=True&loc=Local"
	dbobject, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db = dbobject
}

func GetMySQLDB() *gorm.DB {
	return db
}
