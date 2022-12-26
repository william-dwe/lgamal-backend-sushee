package entity

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	ReviewDescription string
	Rating float64
	OrderedMenuId int
	MenuId int
}