package database

import (
	"go-ambassador/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	DB, err = gorm.Open(mysql.Open("root:root@tcp(ambassador_db:3306)/ambassador"), &gorm.Config{})

	if err != nil {
		panic("Could not connect with the database!" + err.Error())
	}
}

func AutoMigrate() {
	DB.AutoMigrate(models.Product{}, models.Link{}, models.Order{},models.KafkaError{})
}
