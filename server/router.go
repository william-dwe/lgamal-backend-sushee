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
	CartUsecase usecase.CartUsecase
}

func CreateRouter(c RouterConfig) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.JSONifyResult())

	h := handler.New(handler.HandlerConfig{
		UserUsecase:        c.UserUsecase,
		MenuUsecase: c.MenuUsecase,
		CartUsecase: c.CartUsecase,
	})

	v1 := r.Group("/v1")
	v1.GET("/menus", h.ShowMenu)
	v1.GET("/promotions", h.ShowPromotion)
	v1.POST("/login", h.Login)
	v1.POST("/register", h.Register)
	v1.GET("/refresh", h.Refresh)

	user := v1.Group("")
	user.Use(middleware.Authorize)
	user.GET("/users/me", h.ShowUserDetail)
	user.POST("/users/me", h.UpdateUserProfile)

	user.GET("/carts", h.ShowCart)
	user.POST("/carts", h.AddCart)
	
	r.NoRoute(h.NotFoundHandler)

	return r
}
