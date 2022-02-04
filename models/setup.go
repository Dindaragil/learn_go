package models

import (
	"fmt"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	var err error

	Dsn := "host=localhost port=5432 user=postgres dbname=learn_go password=Mieayamkantin4 sslmode=disable"
	database, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqldb, _ := database.DB()

	err = sqldb.Ping()
	if err != nil {
		log.Fatal("database connected")
	}

	fmt.Println("connected to database")
	return database
	// database.AutoMigrate(&Book{}, &User{})

	// DB = database
}

func InitialMigration() {
	database := ConnectDatabase()
	defer Closedatabase(database)
	database.AutoMigrate(&Book{})
	database.AutoMigrate(&User{})
}

func Closedatabase(database *gorm.DB) {
	sqldb, _ := database.DB()
	sqldb.Close()
}
