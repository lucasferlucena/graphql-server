package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lucasferlucena/graphql-server/db/models"
)

var db *gorm.DB //database

//Connect to the database
func Connect() *gorm.DB {
	db, err := gorm.Open("postgres", "user=lucas password=123 dbname=postgres sslmode=disable")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("Connect to postgres database")
	db.Debug().AutoMigrate(&models.User{}, &models.Video{})

	return db
}

//GetDB varieable
func GetDB() *gorm.DB {
	return db
}
