package models

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	Dsn := "host=localhost port=5432 user=postgres dbname=learn_go password=Mieayamkantin4 sslmode=disable"
	database, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&Book{})

	DB = database
}
