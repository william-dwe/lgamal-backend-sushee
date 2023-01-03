package entity

import (
	"time"

	"gorm.io/gorm"
)

type Coupon struct {
	ID uint `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	AdminId int  `json:"admin_id"`
	Description string `json:"description"`
	DiscountAmount float64 `json:"discount_amount"`
	QuotaInitial int `json:"quota_initial"`
	QuotaLeft int `json:"quota_left"`
}

type CouponAddReqBody struct {
	Description string `json:"description"`
	DiscountAmount float64 `json:"discount_amount"`
	QuotaInitial int `json:"quota_initial"`
	QuotaLeft int `json:"quota_left"`
}

type CouponEditReqBody struct {
	Description string `json:"description,omitempty"`
	DiscountAmount float64 `json:"discount_amount,omitempty"`
	QuotaInitial int `json:"quota_initial,omitempty"`
	QuotaLeft int `json:"quota_left,omitempty"`
}

type CouponResBody struct {
	Coupons []Coupon `json:"coupons"`
}

type UserCoupon struct {
	ID uint `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	UserId int `json:"user_id"`
	CouponId int `json:"coupon_id"`
	CouponCode string `json:"coupon_code"`
	DiscountAmount float64 `json:"discount_amount"`
}

type UserCouponAddReqBody struct {
	UserId int `json:"user_id"`
	CouponId int `json:"coupon_id"`
}

type UserCouponResBody struct {
	UserCoupons []UserCoupon `json:"user_coupons"`
}
