package entity

import (
	"time"

	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type Cart struct {
	ID uint `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	UserId int `json:"user_id"`
	MenuId *int `json:"menu_id"`
	Menu Menu `json:"menu"`
	PromotionId *int `json:"promotion_id"`
	Quantity int `json:"quantity"`
	MenuOption pgtype.JSON `json:"menu_option"`
	IsOrdered bool `json:"is_ordered"`
	PromotionPrice *float64 `json:"promotion_price"`
}

type CartResBody struct {
	ID uint `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	UserId int `json:"user_id"`
	MenuId *int `json:"menu_id"`
	PromotionId *int `json:"promotion_id"`
	Quantity int `json:"quantity"`
	MenuOption pgtype.JSON `json:"menu_option"`
	IsOrdered bool `json:"is_ordered"`
	PromotionPrice *float64 `json:"promotion_price"`
}

type CartReqBody struct {
	MenuId *int `json:"menu_id,omitempty"`
	PromotionId *int `json:"promotion_id,omitempty"`
	Quantity int `json:"quantity"`
	MenuOption pgtype.JSON `json:"menu_option,omitempty"`
}

type CartEditDetailsReqBody struct {
	Quantity int `json:"quantity,omitempty"`
	MenuOption pgtype.JSON `json:"menu_option,omitempty"`
}

type CartsResBody struct {
	Carts []Cart `json:"carts"`
}