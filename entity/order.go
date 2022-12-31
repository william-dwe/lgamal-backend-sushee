package entity

import (
	"time"

	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type Order struct {
	ID uint `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	UserId int `json:"user_id"`
	OrderDate time.Time `json:"order_date"`
	CouponId *int `json:"coupon_id"`
	PaymentOptionId int `json:"payment_option_id"`
	OrderedMenus []OrderedMenu `json:"ordered_menus"`
}

type OrderedMenu struct {
	ID uint `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	OrderId int `json:"order_id"`
	MenuId *int `json:"menu_id"`
	PromotionId *int `json:"promotion_id"`
	Quantity int `json:"quantity"`
	MenuOption pgtype.JSON `json:"menu_option"`
}

type OrderReqBody struct {
	CartIdList []int `json:"cart_id_list"`
	PaymentOptionId int `json:"payment_option_id"`
	CouponCode string `json:"coupon_code,omitempty"`
}

type DeliveryOrder struct {
	ID uint `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	OrderId int `json:"order_id"`
	DeliveredAt time.Time `json:"delivered_at"`
}