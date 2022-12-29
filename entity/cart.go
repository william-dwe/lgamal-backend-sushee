package entity

import (
	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId int
	MenuId *int
	Menu Menu
	PromotionId *int
	Quantity int
	MenuOption pgtype.JSON
	IsOrdered bool
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

type CartResBody struct {
	Carts []Cart `json:"carts"`
}