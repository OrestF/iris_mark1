package configs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"../models"
)

func ConfigureDB() *gorm.DB {
	db := setupDBConnection()
	defer db.Close()

	autoMigrateTables(db)
	return db
}

func setupDBConnection() *gorm.DB {
	db, err := gorm.Open("sqlite3", "/home/orest/Training/go/iris_mark1/tmp/gorm.db")

	if err != nil {
		panic("DB error")
	}
	return db
}

func autoMigrateTables(db *gorm.DB) {
	db.AutoMigrate(models.User{})
}
