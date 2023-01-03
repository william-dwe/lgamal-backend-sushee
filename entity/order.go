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
	GrossAmount float64 `json:"gross_amount"` 
	DiscountAmount float64 `json:"discount_amount"`
	NetAmount float64 `json:"net_amount"`
}

type OrderQuery struct {
	Search string
	SortBy string
	FilterByCategory string
	Sort   string
	Limit  int
	Page   int
}

type OrderedMenu struct {
	ID uint `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	OrderId int `json:"order_id"`
	MenuId *int `json:"menu_id"`
	Menu Menu `json:"menu"`
	PromotionId *int `json:"promotion_id"`
	Quantity int `json:"quantity"`
	MenuOption pgtype.JSON `json:"menu_option"`
	Review Review `json:"review"`
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