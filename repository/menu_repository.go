package repository

import (
	"final-project-backend/entity"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MenuRepository interface {
	GetMenu(entity.MenuQuery) (*[]entity.Menu, error)
	GetMenuCount(q entity.MenuQuery) (int, error)
	GetPromotionMenu() (*[]entity.Promotion, error)
	GetAndValidatePromoMenu(menuId, promoId int) (*entity.PromoMenu, error)
	AddMenu(newMenu *entity.Menu) (*entity.Menu, error)
	GetMenuByMenuId(menuId int) (*entity.Menu, error)
	UpdateMenuByMenuId(menuId int, newMenu *entity.Menu) (error)
	DeleteMenuByMenuId(menuId int) (error)
	GetMenuDetailByMenuId(menuId int) (*entity.Menu, error)
}

type MenuRepositoryImpl struct {
	db *gorm.DB
}

type MenuRepositoryConfig struct {
	DB *gorm.DB
}

func NewMenuRepository(c MenuRepositoryConfig) MenuRepository {
	return &MenuRepositoryImpl{
		db: c.DB,
	}
}

func (r *MenuRepositoryImpl) GetMenuCount(q entity.MenuQuery) (int, error) {
	var rows int64

	menuCategorySQ := r.db.
		Select("id").
		Where("category_name IN (?)", strings.Split(q.FilterByCategory, ",")).
		Or("'' = ?", q.FilterByCategory).
		Table("categories")
	query := r.db.
		Joins("menus").
		Where("category_id IN (?)", menuCategorySQ).
		Table("menus")
	query.Count(&rows)
	err := query.Error
	return int(rows), err
}

func (r *MenuRepositoryImpl) GetMenu(q entity.MenuQuery) (*[]entity.Menu, error) {
	var menus []entity.Menu

	menuCategorySQ := r.db.
		Select("id").
		Where("category_name IN (?)", strings.Split(q.FilterByCategory, ",")).
		Or("'' = ?", q.FilterByCategory).
		Table("categories")
	menuSQ := r.db.
		Joins("menus").
		Where("category_id IN (?)", menuCategorySQ).
		Table("menus")
	query := r.db.
		Table("(?) as th", menuSQ).
		Where("menu_name ilike ?", "%"+q.Search+"%").
		Order(clause.OrderByColumn{
			Column: clause.Column{
				Name: q.SortBy,
			},
			Desc: strings.ToLower(q.Sort) == "desc",
		}).
		Limit(q.Limit).
		Offset(q.Page*q.Limit - q.Limit).
		Find(&menus)
	err := query.Error
	return &menus, err
}

func (r *MenuRepositoryImpl) GetPromotionMenu() (*[]entity.Promotion, error) {
	var promotions []entity.Promotion
	err := r.db.
		Model(&entity.Promotion{}).
		Preload("PromoMenus.Menu").
		Where("? between started_at and expired_at", time.Now()).
		Find(&promotions).Error
	return &promotions, err
}

func (r *MenuRepositoryImpl) GetAndValidatePromoMenu(menuId, promoId int) (*entity.PromoMenu, error) {
	var c entity.PromoMenu
	err := r.db.
		Where("menu_id = ? AND promotion_id = ?", menuId, promoId).
		Find(&c).
		Error
	return &c, err
}

func (r *MenuRepositoryImpl) AddMenu(newMenu *entity.Menu) (*entity.Menu, error) {
	err := r.db.
		Create(newMenu).
		Error
	return newMenu, err
}

func (r *MenuRepositoryImpl) GetMenuByMenuId(menuId int) (*entity.Menu, error) {
	var menu entity.Menu

	err := r.db.
		Where("id = ?", menuId).
		First(&menu).Error
	return &menu, err
}

func (r *MenuRepositoryImpl) UpdateMenuByMenuId(menuId int, newMenu *entity.Menu) (error) {
	err := r.db.
		Where("id = ?", menuId).
		Updates(newMenu).
		Error
	return err
}

func (r *MenuRepositoryImpl) DeleteMenuByMenuId(menuId int) (error) {
	var menu entity.Menu

	query := r.db.
		Where("id = (?)", menuId).
		Delete(&menu)

	err := query.Error
	return err
}

func (r *MenuRepositoryImpl) GetMenuDetailByMenuId(menuId int) (*entity.Menu, error) {
	var m entity.Menu

	q := r.db.
		Preload("Reviews").
		Where("id in (?)", menuId).
		Find(&m)
	return &m, q.Error
}