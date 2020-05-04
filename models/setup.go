package models

import (
	"os"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/jinzhu/gorm/dialects/postgres"

)

func SetupModels() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	fmt.Println(os.ExpandEnv("host=${HOST} user=${USER} dbname=${DBNAME} sslmode=disable password=${PASSWORD}"))
	db, err := gorm.Open("postgres", os.ExpandEnv("host=${HOST} user=${USER} dbname=${DBNAME} sslmode=disable password=${PASSWORD}"))

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Spool{})

	return db
}

// func SetupModels() *gorm.DB {
// 	db, err := gorm.Open("sqlite3", "spools.db")

// 	if err != nil {
// 		panic("Failed to connect to database!")
// 	}

// 	db.AutoMigrate(&Spool{})

// 	return db
// }
