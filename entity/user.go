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
	AccessToken string `json:"accessToken"`
}

type UserRegisterReqBody struct {
	FullName  string `json:"full_name" binding:"required"`
	Phone string `gorm:"unique"`
	Email     string `gorm:"unique" binding:"required"`
	Username string `gorm:"unique" binding:"required"`
	Password  string `binding:"required"`
}

type UserRegisterResBody struct {
	FullName  string
	Phone string
	Email     string
	Username string
	Password  string 
	RegisterDate time.Time
	ProfilePicture string
	PlayAttempt int
	RoleId int
}

type UserContext struct {
	Username string
	FullName  string
	Email     string
	Phone string
	ProfilePicture string
	PlayAttempt int
	RoleId int
}

type UserEditDetailsReqBody struct {
	FullName  string 
	Phone 	string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Password  string
	ProfilePicture string 
}