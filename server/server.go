package server

import (
	"fmt"

	"final-project-backend/db"
	"final-project-backend/repository"
	"final-project-backend/usecase"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	userRepo := repository.NewUserRepository(repository.UserRepositoryConfig{
		DB: db.Get(),
	})
	menuRepo := repository.NewMenuRepository(repository.MenuRepositoryConfig{
		DB: db.Get(),
	})
	cartRepo := repository.NewCartRepository(repository.CartRepositoryConfig{
		DB: db.Get(),
	})
	orderRepo := repository.NewOrderRepository(repository.OrderRepositoryConfig{
		DB: db.Get(),
	})
	

	userUsecase := usecase.NewUserUsecase(usecase.UserUsecaseConfig{
		UserRepository:   userRepo,
	})
	menuUsecase := usecase.NewMenuUsecase(usecase.MenuUsecaseConfig{
		MenuRepository:   menuRepo,
	})
	cartUsecase := usecase.NewCartUsecase(usecase.CartUsecaseConfig{
		CartRepository:   cartRepo,
		UserRepository:   userRepo,
		MenuRepository: menuRepo,
	})
	orderUsecase := usecase.NewOrderUsecase(usecase.OrderUsecaseConfig{
		OrderRepository:   orderRepo,
		UserRepository: userRepo,
		CartRepository: cartRepo,
	})


	r := CreateRouter(RouterConfig{
		UserUsecase:        userUsecase,
		MenuUsecase: 		menuUsecase,
		CartUsecase: cartUsecase,
		OrderUsecase: orderUsecase,
	})
	return r
}

func Init() {
	r := initRouter()
	err := r.Run()
	if err != nil {
		fmt.Println("error while running server", err)
		return
	}
}
