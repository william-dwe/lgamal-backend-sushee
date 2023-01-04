package usecase

import (
	"final-project-backend/entity"
	"final-project-backend/errorlist"
	"final-project-backend/repository"
	"fmt"
	"time"
)

type OrderUsecase interface {
	GetPaymentOption() (*[]entity.PaymentOption, error)
	AddOrder(username string, reqBody *entity.OrderReqBody) (*entity.Order, error)
	GetOrderStatus(oq *entity.OrderStatusQuery) (*[]entity.Order, error)
	UpdateOrderStatus(reqBody *entity.OrderStatusUpdateReqBody) (*entity.Order, error)
	GetOrderHistory(username string, oq *entity.OrderHistoryQuery) (*[]entity.Order, error)
	AddReview(username string, r *entity.ReviewAddReqBody) (*entity.Review, error)
}

type orderUsecaseImpl struct {
	orderRepository   repository.OrderRepository
	userRepository repository.UserRepository
	cartRepository repository.CartRepository
	couponRepository repository.CouponRepository
}

type OrderUsecaseConfig struct {
	OrderRepository   repository.OrderRepository
	UserRepository repository.UserRepository
	CartRepository repository.CartRepository
	CouponRepository repository.CouponRepository

}

func NewOrderUsecase(c OrderUsecaseConfig) OrderUsecase {
	return &orderUsecaseImpl{
		orderRepository:   c.OrderRepository,
		userRepository: c.UserRepository,
		cartRepository: c.CartRepository,
		couponRepository: c.CouponRepository,
	}
}

func validateCartOwnershipAndAvailability(cart *entity.Cart, user *entity.User) error {
	if cart.UserId != int(user.ID) {
		return errorlist.UnauthorizedError()
	} 
	if cart.IsOrdered {
		return errorlist.BadRequestError("the cart has been ordered before", "INVALID_CART")
	} 

	return nil
}


func (u *orderUsecaseImpl) GetPaymentOption() (*[]entity.PaymentOption, error) {
	paymentOptions, err :=  u.orderRepository.GetPaymentOption()
	if err != nil {
		return nil, errorlist.InternalServerError()
	}
	
	return paymentOptions, nil
}

func (u *orderUsecaseImpl) AddOrder(username string, reqBody *entity.OrderReqBody) (*entity.Order, error) {
	if len(reqBody.CartIdList) == 0 {
		return nil, errorlist.BadRequestError("Please order something", "INVALID_ORDER")
	}
	user, err := u.userRepository.GetUserByEmailOrUsername(username)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}

	carts, err :=  u.cartRepository.GetCartByCartIds(reqBody.CartIdList)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}
	if len(*carts) != len(reqBody.CartIdList) {
		return nil, errorlist.BadRequestError("Some cart IDs are not found", "INVALID_CART_IDS")
	}
	for _, cart := range *carts {
		err = validateCartOwnershipAndAvailability(&cart, user)
		if err != nil {
			return nil, err
		}
	}

	newOrder := entity.Order{
		UserId: int(user.ID),
		OrderDate: time.Now(),
		PaymentOptionId: reqBody.PaymentOptionId,
		Status: "Payment",
	}

	var couponId int
	if reqBody.CouponCode != "" {
		coupon, err := u.couponRepository.GetUserCouponByCouponCode(int(user.ID), reqBody.CouponCode)
		if err != nil {
			return nil, errorlist.InternalServerError()
		}
		couponId = int(coupon.ID) 
		newOrder.CouponId = &couponId
		newOrder.DiscountAmount = coupon.DiscountAmount
	}
	
	totalPrice, err :=  u.cartRepository.GetCartTotalPriceByCartIds(reqBody.CartIdList)
	if err != nil {
		return nil, err
	}
	newOrder.GrossAmount = totalPrice
	newOrder.NetAmount = newOrder.GrossAmount-newOrder.DiscountAmount

	order, err := u.orderRepository.AddOrder(&newOrder)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}


	var newOrderedMenus []entity.OrderedMenu
	for _, cart := range *carts {
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

	
	u.cartRepository.UpdateCartByCartIds(reqBody.CartIdList, &entity.Cart{
		IsOrdered: true,
	})

	if reqBody.CouponCode != "" {
		u.couponRepository.DeleteCouponById(*newOrder.CouponId)
	}

	return order, nil
}

func (u *orderUsecaseImpl) GetOrderStatus(oq *entity.OrderStatusQuery) (*[]entity.Order, error) {
	orders, err := u.orderRepository.GetOrderStatus(*oq)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}

	return orders, nil
}

func (u *orderUsecaseImpl) UpdateOrderStatus(reqBody *entity.OrderStatusUpdateReqBody) (*entity.Order, error) {
	orderWithNewStatus := entity.Order{
		Status: reqBody.Status,
	}
	
	err := u.orderRepository.UpdateOrderByOrderId(reqBody.ID, &orderWithNewStatus)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}

	order, err := u.orderRepository.GetOrderById(reqBody.ID)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}

	return order, nil
}



func (u *orderUsecaseImpl) GetOrderHistory(username string, oq *entity.OrderHistoryQuery) (*[]entity.Order, error) {
	user, err := u.userRepository.GetUserByEmailOrUsername(username)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}
	
	orders, err := u.orderRepository.GetOrderHistory(int(user.ID), *oq)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}

	return orders, nil
}

func (u *orderUsecaseImpl) AddReview(username string, r *entity.ReviewAddReqBody) (*entity.Review, error) {
	user, err := u.userRepository.GetUserByEmailOrUsername(username)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}

	orderedMenu, err :=  u.orderRepository.GetOrderedMenuById(r.OrderedMenuId)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}

	order, err := u.orderRepository.GetOrderById(orderedMenu.OrderId)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}
	if order.UserId != int(user.ID) {
		return nil, errorlist.UnauthorizedError()
	}

	newReview := entity.Review {
		ReviewDescription: r.ReviewDescription,
		Rating: r.Rating,
		OrderedMenuId: r.OrderedMenuId,
		MenuId: *orderedMenu.MenuId,
	}

	review, err := u.orderRepository.AddReview(&newReview)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}

	err = u.orderRepository.UpdateAvgReviewScoreByMenuId(newReview.MenuId)
	if err != nil {
		fmt.Println("ERR:", err)
		return nil, errorlist.InternalServerError()
	}

	return review, nil
}
