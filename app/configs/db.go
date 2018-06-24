package configs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Email string
}

func ConfigureDB() {
	db := setupDBConnection()
	defer db.Close()

	autoMigrateTables(db)

	// CREATE, FIND TESTING
	user := User{Email: "user@mail2.com"}
	db.Create(&user)

	var u3 User
	db.First(&u3)
	println("FINDED USER", u3.Email)
}

func setupDBConnection() *gorm.DB {
	db, err := gorm.Open("sqlite3", "/tmp/gorm.db")
	// defer db.Close()

	if err != nil {
		panic("DB error")
	}
	return db
}

func autoMigrateTables(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
