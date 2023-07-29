package database

import (
	"backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var dsn = "root:dinesh@tcp(127.0.0.1:3306)/buzzerquiz"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database server")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
	// conn, err := gorm
}
