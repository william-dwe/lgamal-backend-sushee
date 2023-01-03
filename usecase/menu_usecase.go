package usecase

import (
	"errors"
	"final-project-backend/entity"
	"final-project-backend/errorlist"
	"final-project-backend/repository"

	"gorm.io/gorm"
)

type MenuUsecase interface {
	GetMenu(entity.MenuQuery) (*[]entity.Menu, int, error)
	GetPromotion() (*[]entity.Promotion, error)
	AddMenu(reqBody *entity.MenuAddReqBody) (*entity.Menu, error)
	UpdateMenuByMenuId(menuId int, m *entity.Menu) (*entity.Menu, error)
	DeleteMenuByMenuId(menuId int) (error)
	GetMenuDetailByMenuId(menuId int) (*entity.Menu, error)
}

type menuUsecaseImpl struct {
	menuRepository   repository.MenuRepository
}

type MenuUsecaseConfig struct {
	MenuRepository   repository.MenuRepository
}

func NewMenuUsecase(c MenuUsecaseConfig) MenuUsecase {
	return &menuUsecaseImpl{
		menuRepository:   c.MenuRepository,
	}
}

func (u *menuUsecaseImpl) GetMenu(q entity.MenuQuery) (*[]entity.Menu, int, error) {
	rows, err := u.menuRepository.GetMenuCount(q)
	if err != nil {
		return nil, 0, errorlist.InternalServerError()
	}

	if q.Limit == 0 {
		q.Limit = rows	
	}
	menus, err := u.menuRepository.GetMenu(q)
	if errors.Is(err, gorm.ErrRecordNotFound) || len(*menus) == 0 {
		return nil, 0, errorlist.BadRequestError("no menu found", "NO_MENU_FOUND")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, errorlist.InternalServerError()
	}
	
	maxPage := (rows + q.Limit - 1)/ q.Limit
	return menus, maxPage, nil
}


func (u *menuUsecaseImpl) GetPromotion() (*[]entity.Promotion, error) {
	menus, err := u.menuRepository.GetPromotionMenu()
	if errors.Is(err, gorm.ErrRecordNotFound) || len(*menus) == 0 {
		return nil, errorlist.BadRequestError("no available promotion", "NO_MENU_FOUND")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorlist.InternalServerError()
	}
	
	return menus, nil
}

func (u *menuUsecaseImpl) AddMenu(reqBody *entity.MenuAddReqBody) (*entity.Menu, error) {
	newMenu := entity.Menu{
		MenuName: reqBody.MenuName,
		Price: reqBody.Price,
		MenuPhoto: reqBody.MenuPhoto,
		CategoryId: reqBody.CategoryId,
		Customization: guardNullJSON(reqBody.Customization),
	}

	menu, err := u.menuRepository.AddMenu(&newMenu)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}
	
	return menu, nil
}

func (u *menuUsecaseImpl) UpdateMenuByMenuId(menuId int, m *entity.Menu) (*entity.Menu, error) {
	err := u.menuRepository.UpdateMenuByMenuId(menuId, m)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}

	menu, err := u.menuRepository.GetMenuByMenuId(menuId)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}
	
	return menu, nil
}

func (u *menuUsecaseImpl) DeleteMenuByMenuId(menuId int) (error) {
	err := u.menuRepository.DeleteMenuByMenuId(menuId)
	if err != nil {
		return errorlist.InternalServerError()
	}

	return nil
}

func (u *menuUsecaseImpl) GetMenuDetailByMenuId(menuId int) (*entity.Menu, error) {
	menu, err := u.menuRepository.GetMenuDetailByMenuId(menuId)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}

	return menu, nil
}