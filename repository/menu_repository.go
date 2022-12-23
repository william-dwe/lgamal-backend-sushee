package repository

import (
	"final-project-backend/entity"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MenuRepository interface {
	GetMenu(entity.MenuQuery) (*[]entity.Menu, int, error)
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

func (r *MenuRepositoryImpl) GetMenu(q entity.MenuQuery) (*[]entity.Menu, int, error) {
	var Menus []entity.Menu

	MenuCategorySQ := r.db.
		Select("id").
		Where("category_name IN (?)", strings.Split(q.FilterByCategory, ",")).
		Or("'' = ?", q.FilterByCategory).
		Table("categories")
	MenuSubQuery := r.db.
		Joins("menus").
		Where("category_id IN (?)", MenuCategorySQ).
		Table("menus")
	query := r.db.
		Table("(?) as th", MenuSubQuery).
		Where("menu_name ilike ?", "%"+q.Search+"%").
		Order(clause.OrderByColumn{
			Column: clause.Column{
				Name: q.SortBy,
			},
			Desc: strings.ToLower(q.Sort) == "desc",
		}).
		Limit(q.Limit).
		Offset(q.Page*q.Limit - q.Limit).
		Find(&Menus)
	var rows int64
	MenuSubQuery.Count(&rows)
	err := query.Error
	return &Menus, int(rows), err
}
