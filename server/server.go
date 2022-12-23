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

	userUsecase := usecase.NewUserUsecase(usecase.UserUsecaseConfig{
		UserRepository:   userRepo,
	})
	menuUsecase := usecase.NewMenuUsecase(usecase.MenuUsecaseConfig{
		MenuRepository:   menuRepo,
	})


	r := CreateRouter(RouterConfig{
		UserUsecase:        userUsecase,
		MenuUsecase: 		menuUsecase,
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
