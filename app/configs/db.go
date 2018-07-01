package configs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"../models"
)

func ConfigureDB() *gorm.DB {
	db := setupDBConnection()
	models.DB = db

	autoMigrateTables(db)
	return db
}

func setupDBConnection() *gorm.DB {
	db, err := gorm.Open("sqlite3", "../iris_mark1/tmp/gorm.db")

	if err != nil {
		panic(err)
	}
	return db
}

func autoMigrateTables(db *gorm.DB) {
	db.AutoMigrate(models.User{})
}
