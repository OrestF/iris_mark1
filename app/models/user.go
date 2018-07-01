package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email string
	Name  string
}

type UserModel struct {
}

func (um *UserModel) Delete(id int64) {
	DB.Delete(User{}, "id = ?", id)
}

func (um *UserModel) All() []User {
	var users []User
	DB.Find(&users)
	return users
}

func (um *UserModel) Find(id int64) User {
	var user User
	DB.First(&user, id)
	return user
}

func (um *UserModel) Create(user User) User {
	DB.Create(&user)
	return user
}

func (um *UserModel) Update(id int64, userParamsStruct User) User {
	dbUser := um.Find(id)
	DB.Model(&dbUser).Updates(userParamsStruct)
	return dbUser
}
