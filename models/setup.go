package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("sqlite3", "spools.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Spool{})

	return db
}
