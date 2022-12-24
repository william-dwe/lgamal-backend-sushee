package usecase

import (
	"errors"
	"final-project-backend/entity"
	"final-project-backend/errorlist"
	"final-project-backend/repository"

	"gorm.io/gorm"
)

type CartUsecase interface {
	GetCart(username string) (*[]entity.Cart, error)
	AddCart(username string, c *entity.CartReqBody) (*entity.Cart, error)
}

type cartUsecaseImpl struct {
	cartRepository   repository.CartRepository
	userRepository repository.UserRepository
}

type CartUsecaseConfig struct {
	CartRepository   repository.CartRepository
	UserRepository repository.UserRepository
}

func NewCartUsecase(c CartUsecaseConfig) CartUsecase {
	return &cartUsecaseImpl{
		cartRepository:   c.CartRepository,
		userRepository: c.UserRepository,
	}
}


func (u *cartUsecaseImpl) GetCart(username string) (*[]entity.Cart, error) {
	carts, err := u.cartRepository.GetCartByUsername(username)
	if errors.Is(err, gorm.ErrRecordNotFound) || len(*carts) == 0 {
		return nil, errorlist.BadRequestError("no cart found", "NO_CART_FOUND")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorlist.InternalServerError()
	}
	
	return carts, nil
}


func (u *cartUsecaseImpl) AddCart(username string, c *entity.CartReqBody) (*entity.Cart, error) {
	user, err := u.userRepository.GetUserByEmailOrUsername(username)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}
	
	newCartItem := entity.Cart{
		UserId: int(user.ID),
		MenuId: c.MenuId,
		PromotionId: c.PromotionId,
		Quantity: c.Quantity,
		MenuOption: c.MenuOption,
	}

	cart, err := u.cartRepository.AddItemToCart(&newCartItem)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}
	
	return cart, nil
}
