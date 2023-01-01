package repository

import (
	"final-project-backend/entity"

	"gorm.io/gorm"
)

type OrderRepository interface {
	GetPaymentOption() (*[]entity.PaymentOption, error)
	AddOrder(newOrder *entity.Order) (*entity.Order, error)
	AddOrderedMenu(newOrderedMenus *[]entity.OrderedMenu) (*[]entity.OrderedMenu, error)
	GetOrderHistory(userId int) (*[]entity.Order, error)
}

type OrderRepositoryImpl struct {
	db *gorm.DB
}

type OrderRepositoryConfig struct {
	DB *gorm.DB
}

func NewOrderRepository(c OrderRepositoryConfig) OrderRepository {
	return &OrderRepositoryImpl{
		db: c.DB,
	}
}



func (r *OrderRepositoryImpl) GetPaymentOption() (*[]entity.PaymentOption, error) {
	var payments []entity.PaymentOption
	err := r.db.
		Find(&payments).Error
	return &payments, err
}

func (r *OrderRepositoryImpl) AddOrder(newOrder *entity.Order) (*entity.Order, error) {
	err := r.db.
		Create(newOrder).
		Error
	return newOrder, err
}


func (r *OrderRepositoryImpl) AddOrderedMenu(newOrderedMenus *[]entity.OrderedMenu) (*[]entity.OrderedMenu, error) {
	err := r.db.
		Create(newOrderedMenus).
		Error
	return newOrderedMenus, err
}

func (r *OrderRepositoryImpl) GetOrderHistory(userId int) (*[]entity.Order, error) {
	var o []entity.Order
	q := r.db.
		Preload("OrderedMenus").
		Where("user_id in (?)", userId).
		Find(&o)
	return &o, q.Error
}
