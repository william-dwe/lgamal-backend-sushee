package server

import (
	"final-project-backend/handler"
	"final-project-backend/middleware"
	"final-project-backend/usecase"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	UserUsecase        usecase.UserUsecase
	MenuUsecase 	usecase.MenuUsecase
}

func CreateRouter(c RouterConfig) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.JSONifyResult())

	h := handler.New(handler.HandlerConfig{
		UserUsecase:        c.UserUsecase,
		MenuUsecase: c.MenuUsecase,
	})

	v1 := r.Group("/v1")
	v1.GET("/menu", h.ShowMenu)
	v1.POST("/login", h.Login)
	v1.POST("/register", h.Register)
	v1.GET("/refresh", h.Refresh)

	user := v1.Group("")
	user.Use(middleware.Authorize)
	user.GET("/users/me", h.UserDetail)
	user.POST("/users/me", h.UpdateUserProfile)
	
	r.NoRoute(h.NotFoundHandler)

	return r
}
