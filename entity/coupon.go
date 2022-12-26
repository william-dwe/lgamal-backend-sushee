package entity

import (
	"gorm.io/gorm"
)

type Coupon struct {
	gorm.Model
	AdminId int
	Description string
	CouponDiscountAmount float64
	QuotaInitial int
	QuotaLeft int
}

type UserCoupon struct {
	gorm.Model
	UserId int
	CouponId int
	CouponCode string
}