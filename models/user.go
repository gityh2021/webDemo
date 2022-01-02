package models

import (
	"database/sql"
	"gorm.io/gorm"
)
type User struct {
	gorm.Model
	Age sql.NullInt64
	Name string `gorm:"not null" json:"name" binding:"required"`
	Password string `gorm:"not null" json:"password" binding:"required"`
	Info string `json:"info"`
	StuffNo string `gorm:"not null" json:"stuff_no" binding:"required"`
	Department string `json:"department"`
	Active bool `gorm:"not null" json:"active"`
}

func (user *User) Register() (*sql.Rows, error) {
	result := db.Create(user)
	user.Password = "guess!"
	return result.Rows()
}

func GetUserById(id int) User {
	var user User
	db.First(&user, id)
	user.Password = "guess!"
	return user
}