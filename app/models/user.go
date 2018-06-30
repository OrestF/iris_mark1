package models

import (
	"github.com/jinzhu/gorm"
)

type UserModel struct {
	db gorm.DB
}

type User struct {
	gorm.Model
	Email string
	Name  string
}

func NewUserModel() UserModel {
	db, err := gorm.Open("sqlite3", "/home/orest/Training/go/iris_mark1/tmp/gorm.db")
	if err != nil {
		panic("DB error")
	}
	return UserModel{db: *db}
}

func (um *UserModel) Delete(id int64) {
	um.db.Delete(User{}, "id = ?", id)
}

func (um *UserModel) All() []User {
	var users []User
	um.db.Find(&users)
	return users
}

func (um *UserModel) Find(id int64) User {
	var user User
	um.db.First(&user, id)
	return user
}

func (um *UserModel) Create(user User) User {
	um.db.Create(&user)
	return user
}
