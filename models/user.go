package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password" gorm:"not null"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
}
