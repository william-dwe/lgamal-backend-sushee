package entity

import (
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	MenuName string
	AvgRating int
	NumberOfFavorites int
	Price float64
	MenuPhoto string
	CategoryId int
}

type MenuQuery struct {
	Search string
	SortBy string
	FilterByCategory string
	Sort   string
	Limit  int
	Page   int
}

type MenuResBody struct {
	Menus []Menu `json:"menus"`
	CurrentPage int `json:"current_page"`
	MaxPage int `json:"max_page"`
}

type MenuCategory struct {
	gorm.Model
	CategoryName string
}