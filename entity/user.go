package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName  string
	Phone string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Username string `gorm:"unique"`
	Password  string 
	RegisterDate time.Time
	ProfilePicture string
	PlayAttempt int
	RoleId int
}

type UserLoginReqBody struct {
	Identifier    string `json:"identifier" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResBody struct {
	AccessToken string `json:"access_token"`
}
