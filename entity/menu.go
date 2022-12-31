package entity

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	ID uint `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	MenuName string `json:"menu_name"`
	AvgRating int `json:"avg_rating"`
	NumberOfFavorites int `json:"number_of_favorites"`
	Price float64 `json:"price"`
	MenuPhoto string `json:"menu_photo"`
	CategoryId int `json:"category_id"`
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

type PromoMenu struct {
	ID uint `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	PromotionId int `json:"promotion_id"`
	MenuId int `json:"menu_id"`
	Menu Menu `json:"menu" gorm:"foreignKey:MenuId;references:ID"`
	PromotionPrice float64 `json:"promotion_price"`
}

type Promotion struct {
	ID uint `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	AdminId int `json:"admin_id"`
	Name string `json:"name"`
	Description string `json:"description"`
	PromotionPhoto string `json:"promotion_photo"`
	DiscountRate float64 `json:"discount_rate"`
	StartAt time.Time `json:"start_at"`
	ExpiredAt time.Time `json:"expired_at"`
	PromoMenus []PromoMenu `json:"promo_menus" gorm:"foreignKey:PromotionId;references:ID"`
}

type PromotionResBody struct {
	Promotions []Promotion `json:"promotions"`
}