package entity

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	RefreshToken string `gorm:"unique"`
	UserId int 
	ExpiredAt time.Time
}