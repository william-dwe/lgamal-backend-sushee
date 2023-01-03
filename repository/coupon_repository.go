package repository

import (
	"final-project-backend/entity"
	"fmt"

	"gorm.io/gorm"
)

type CouponRepository interface {
	AddCoupon(coupon *entity.Coupon) (*entity.Coupon, error)
	GetCoupon() (*[]entity.Coupon, error)
	GetCouponById(couponId int) (*entity.Coupon, error)
	UpdateCouponById(couponId int, newCoupon *entity.Coupon) (error)
	DeleteCouponById(couponId int) (*entity.Coupon, error)
	AddUserCoupon(userCoupon *entity.UserCoupon) (*entity.UserCoupon, error)
	GetUserCouponByUsername(username string) (*[]entity.UserCoupon, int, error)
	GetUserCouponByCouponCode(userId int, couponCode string) (*entity.UserCoupon, error)
}

type CouponRepositoryImpl struct {
	db *gorm.DB
}

type CouponRepositoryConfig struct {
	DB *gorm.DB
}

func NewCouponRepository(c CouponRepositoryConfig) CouponRepository {
	return &CouponRepositoryImpl{
		db: c.DB,
	}
}

func (r *CouponRepositoryImpl) AddCoupon(coupon *entity.Coupon) (*entity.Coupon, error) {
	err := r.db.
		Create(&coupon).Error
	return coupon, err
}

func (r *CouponRepositoryImpl) GetCoupon() (*[]entity.Coupon, error) {
	var coupon []entity.Coupon
	err := r.db.
		Find(&coupon).Error
	return &coupon, err
}

func (r *CouponRepositoryImpl) GetCouponById(couponId int) (*entity.Coupon, error) {
	var coupon entity.Coupon

	err := r.db.
		Where("id = (?)", couponId).Debug().
		First(&coupon).Error
	return &coupon, err
}

func (r *CouponRepositoryImpl) UpdateCouponById(couponId int, newCoupon *entity.Coupon) (error) {
	err := r.db.
		Where("id = (?)", couponId).
		Updates(newCoupon).
		Debug().Error
	return err
}

func (r *CouponRepositoryImpl) DeleteCouponById(couponId int) (*entity.Coupon, error) {
	var coupon entity.Coupon
	err := r.db.
		Where("id = (?)",couponId).
		Delete(&coupon).Error
	return &coupon, err
}

func (r *CouponRepositoryImpl) AddUserCoupon(userCoupon *entity.UserCoupon) (*entity.UserCoupon, error) {
	err := r.db.
		Create(&userCoupon).
		Error
	return userCoupon, err
}

func (r *CouponRepositoryImpl) GetUserCouponByUsername(username string) (*[]entity.UserCoupon, int, error) {
	var coupon []entity.UserCoupon
	userSQ := r.db.
		Table("users").
		Select("id").
		Where("username = (?)", username)
	q := r.db.
		Where("user_id in (?)", userSQ).
		Find(&coupon)
	return &coupon, int(q.RowsAffected), q.Error
}

func (r *CouponRepositoryImpl) GetUserCouponByCouponCode(userId int, couponCode string) (*entity.UserCoupon, error) {
	var coupon entity.UserCoupon
	fmt.Println("DEBUG - userid", userId)
	fmt.Println("DEBUG - coupon", couponCode)
	q := r.db.
		Where("user_id in (?) AND coupon_code = (?)", userId, couponCode).
		First(&coupon)
	return &coupon, q.Error
}
