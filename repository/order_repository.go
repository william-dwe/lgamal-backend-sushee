package repository

import (
	"final-project-backend/entity"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderRepository interface {
	GetPaymentOption() (*[]entity.PaymentOption, error)
	AddOrder(newOrder *entity.Order) (*entity.Order, error)
	AddOrderedMenu(newOrderedMenus *[]entity.OrderedMenu) (*[]entity.OrderedMenu, error)
	GetOrderStatus(oq entity.OrderStatusQuery) (*[]entity.Order, error)
	GetOrderHistory(userId int, oq entity.OrderHistoryQuery) (*[]entity.Order, error)
	GetOrderById(orderId int) (*entity.Order, error)
	UpdateOrderByOrderId(orderId int, newOrderStatus *entity.Order) (error)
	GetOrderedMenuById(orderedMenuId int) (*entity.OrderedMenu, error)
	AddReview(review *entity.Review) (*entity.Review, error)
	UpdateAvgReviewScoreByMenuId(MenuId int) (error)
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


func (r *OrderRepositoryImpl) GetOrderStatus(oq entity.OrderStatusQuery) (*[]entity.Order, error) {
	var o []entity.Order

	sqSelectedMenuOrder := r.db.
		Table("orders o").
		Select("user_id, order_id, menu_id").
		Joins("join ordered_menus om ON o.id = om.order_id").
		Where("status in (?)", oq.FilterByStatus)

	sqSelectedMenu := r.db.
		Table("menus m").
		Select("order_id").
		Joins("join (?) sm  ON sm.menu_id = m.id", sqSelectedMenuOrder).
		Where("menu_name ilike (?)", "%"+oq.Search+"%")
		
	q := r.db.
		Preload("OrderedMenus").
		Preload("OrderedMenus.Menu").
		Preload("OrderedMenus.Review").
		Where("id in (?)", sqSelectedMenu).
		Order(clause.OrderByColumn{
			Column: clause.Column{
				Name: oq.SortBy,
			},
			Desc: strings.ToLower(oq.Sort) == "desc",
		}).
		Limit(oq.Limit).
		Offset(oq.Page*oq.Limit - oq.Limit).
		Find(&o)
	return &o, q.Error
}

func (r *OrderRepositoryImpl) GetOrderHistory(userId int, oq entity.OrderHistoryQuery) (*[]entity.Order, error) {
	var o []entity.Order

	sqSelectedMenuOrder := r.db.
		Table("orders o").
		Select("user_id, order_id, menu_id").
		Joins("join ordered_menus om ON o.id = om.order_id").
		Where("user_id in (?)", userId)

	sqSelectedMenu := r.db.
		Table("menus m").
		Select("order_id").
		Joins("join (?) sm  ON sm.menu_id = m.id", sqSelectedMenuOrder).
		Where("menu_name ilike (?)", "%"+oq.Search+"%")
		
	q := r.db.
		Preload("OrderedMenus").
		Preload("OrderedMenus.Menu").
		Preload("OrderedMenus.Review").
		Where("id in (?)", sqSelectedMenu).
		Order(clause.OrderByColumn{
			Column: clause.Column{
				Name: oq.SortBy,
			},
			Desc: strings.ToLower(oq.Sort) == "desc",
		}).
		Limit(oq.Limit).
		Offset(oq.Page*oq.Limit - oq.Limit).
		Find(&o)
	return &o, q.Error
}


func (r *OrderRepositoryImpl) GetOrderById(orderId int) (*entity.Order, error) {
	var o entity.Order

	err := r.db.
	Where("id = (?)", orderId).
	First(&o).
	Error
	return &o, err
}

func (r *OrderRepositoryImpl) UpdateOrderByOrderId(orderId int, newOrderStatus *entity.Order) (error) {
	err := r.db.
		Where("id = ?", orderId).
		Updates(newOrderStatus).
		Error
	return err
}


func (r *OrderRepositoryImpl) GetOrderedMenuById(orderedMenuId int) (*entity.OrderedMenu, error) {
	var o entity.OrderedMenu

	err := r.db.
	Where("id = (?)", orderedMenuId).
	First(&o).
	Error
	return &o, err
}

func (r *OrderRepositoryImpl) AddReview(review *entity.Review) (*entity.Review, error) {
	err := r.db.
	Create(review).
	Error
return review, err
}

func (r *OrderRepositoryImpl) UpdateAvgReviewScoreByMenuId(MenuId int) (error) {
	sqAvgReview := r.db.
	Table("reviews").
	Select("AVG(rating)")

	err := r.db.
	Table("Menus").
	Where("id = (?)", MenuId).
	Update("avg_rating", sqAvgReview).
	Error
return err
}