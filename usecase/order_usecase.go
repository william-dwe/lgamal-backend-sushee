package usecase

import (
	"errors"
	"final-project-backend/entity"
	"final-project-backend/errorlist"
	"final-project-backend/repository"
	"time"

	"gorm.io/gorm"
)

type OrderUsecase interface {
	GetPaymentOption() (*[]entity.PaymentOption, error)
	GetCoupon() (*[]entity.Coupon, error)
	GetUserCouponByUsername(username string) (*[]entity.UserCoupon, error)
	AddOrder(username string, reqBody *entity.OrderReqBody) (*entity.Order, error)
	GetOrderHistory(username string) (*[]entity.Order, error)
}

type orderUsecaseImpl struct {
	orderRepository   repository.OrderRepository
	userRepository repository.UserRepository
	cartRepository repository.CartRepository
}

type OrderUsecaseConfig struct {
	OrderRepository   repository.OrderRepository
	UserRepository repository.UserRepository
	CartRepository repository.CartRepository
}

func NewOrderUsecase(c OrderUsecaseConfig) OrderUsecase {
	return &orderUsecaseImpl{
		orderRepository:   c.OrderRepository,
		userRepository: c.UserRepository,
		cartRepository: c.CartRepository,
	}
}


func (u *orderUsecaseImpl) GetPaymentOption() (*[]entity.PaymentOption, error) {
	paymentOptions, err :=  u.orderRepository.GetPaymentOption()
	if err != nil {
		return nil, errorlist.InternalServerError()
	}
	
	return paymentOptions, nil
}


func (u *orderUsecaseImpl) GetCoupon() (*[]entity.Coupon, error) {
	coupons, err :=  u.orderRepository.GetCoupon()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorlist.BadRequestError("no coupon available", "NO_COUPON_EXIST")
	}
	if err != nil {
		return nil, errorlist.InternalServerError()
	}
	
	return coupons, nil
}


func (u *orderUsecaseImpl) GetUserCouponByUsername(username string) (*[]entity.UserCoupon, error) {
	userCoupons, r, err :=  u.orderRepository.GetUserCouponByUsername(username)
	if errors.Is(err, gorm.ErrRecordNotFound) || r == 0 {
		return nil, errorlist.BadRequestError("no coupon available", "NO_COUPON_EXIST")
	}
	if err != nil {
		return nil, errorlist.InternalServerError()
	}
	
	return userCoupons, nil
}



func (u *orderUsecaseImpl) AddOrder(username string, reqBody *entity.OrderReqBody) (*entity.Order, error) {
	user, err := u.userRepository.GetUserByEmailOrUsername(username)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}


	newOrder := entity.Order{
		UserId: int(user.ID),
		OrderDate: time.Now(),
		PaymentOptionId: reqBody.PaymentOptionId,
	}

	if reqBody.CouponCode != "" {
		coupon, r, err := u.orderRepository.GetUserCouponByCouponCode(int(user.ID), reqBody.CouponCode)
		if errors.Is(err, gorm.ErrRecordNotFound) || r == 0{
			return nil, errorlist.BadRequestError("coupon code invalid", "INVALID_COUPON_CODE")
		}
		if err != nil {
			return nil, errorlist.InternalServerError()
		}
		couponId := int(coupon.ID) 
		newOrder.CouponId = &couponId
	}

	order, err := u.orderRepository.AddOrder(&newOrder)

	if err != nil {
		return nil, errorlist.InternalServerError()
	}


	var newOrderedMenus []entity.OrderedMenu
	for _, c := range reqBody.CartIdList {
		cart, err := u.cartRepository.GetCartByCartId(c)
		if err != nil {
			return nil, errorlist.BadRequestError("Cart IDs not found", "INVALID_CART_IDS")
		}

		o := entity.OrderedMenu{
			OrderId: int(order.ID),
			Quantity: cart.Quantity,
			MenuOption: cart.MenuOption,
		}
		if cart.MenuId != nil {
			o.MenuId = cart.MenuId
		}
		if cart.PromotionId != nil {
			o.PromotionId = cart.PromotionId

		}
		newOrderedMenus = append(newOrderedMenus, o)
	}

	_, err = u.orderRepository.AddOrderedMenu(&newOrderedMenus)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}

	for _, id := range reqBody.CartIdList {
		u.cartRepository.DeleteCartByCartId(id)
	}

	// todo: delivery

	return order, nil
}


func (u *orderUsecaseImpl) GetOrderHistory(username string) (*[]entity.Order, error) {
	user, err := u.userRepository.GetUserByEmailOrUsername(username)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}

	order, err := u.orderRepository.GetOrderHistory(int(user.ID))
	if err != nil {
		return nil, errorlist.InternalServerError()
	}

	return order, nil
}

// todo: set delivery status
// plan: Prepared --> Sending --> Received
// Details:
// P -> default val
// S -> admin confirmed
// S -> admin confirmed