package repository

import (
	"final-project-backend/entity"

	"gorm.io/gorm"
)

type CartRepository interface {
	AddItemToCart(c *entity.Cart) (*entity.Cart, error)
	GetCartByUsername(username string) (*[]entity.Cart, error)
	// DeleteCart(username string) (error)
	// DeleteCartByCartId(username string, cartId int) (error)
	// UpdateCartByCartId(username string, cartId int, updatePremises *entity.Cart) (*entity.Cart, error)
}

type CartRepositoryImpl struct {
	db *gorm.DB
}

type CartRepositoryConfig struct {
	DB *gorm.DB
}

func NewCartRepository(c CartRepositoryConfig) CartRepository {
	return &CartRepositoryImpl{
		db: c.DB,
	}
}

func (r *CartRepositoryImpl) AddItemToCart(c *entity.Cart) (*entity.Cart, error) {
	err := r.db.Create(&c).Error
	return c, err
}

func (r *CartRepositoryImpl) GetCartByUsername(username string) (*[]entity.Cart, error) {
	var carts []entity.Cart

	userSQ := r.db.
		Select("id").
		Where("username = (?)", username).
		Table("users")
	menuSubQuery := r.db.
		Where("is_ordered != (?)", true).
		Table("carts")
	query := r.db.
		Table("(?) as th", menuSubQuery).
		Where("user_id = (?)", userSQ).
		Find(&carts)
		
	err := query.Error
	return &carts, err
}
