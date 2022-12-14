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

	userUsecase := usecase.NewUserUsecase(usecase.UserUsecaseConfig{
		UserRepository:   userRepo,
	})

	r := CreateRouter(RouterConfig{
		UserUsecase:        userUsecase,
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
