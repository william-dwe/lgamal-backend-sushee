package entity

import (
	"time"

	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserId int
	OrderDate time.Time
	CouponId *int
	PaymentOptionId int
	OrderedMenus []OrderedMenu
}

type OrderedMenu struct {
	gorm.Model
	OrderId int
	MenuId *int
	PromotionId *int
	Quantity int
	MenuOption pgtype.JSON
}

type OrderReqBody struct {
	CartIdList []int `json:"cart_id_list"`
	PaymentOptionId int `json:"payment_option_id"`
	CouponCode string `json:"coupon_code,omitempty"`
}

type DeliveryOrder struct {
	gorm.Model
	OrderId int
	DeliveredAt time.Time
}