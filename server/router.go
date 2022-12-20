package server

import (
	"final-project-backend/handler"
	"final-project-backend/middleware"
	"final-project-backend/usecase"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	UserUsecase        usecase.UserUsecase
}

func CreateRouter(c RouterConfig) *gin.Engine {
	r := gin.Default()

	h := handler.New(handler.HandlerConfig{
		UserUsecase:        c.UserUsecase,
	})

	r.Use(middleware.JSONifyResult())
	r.POST("/login", h.Login)
	r.POST("/register", h.Register)
	r.GET("/refresh", h.Refresh)

	r.Use(middleware.Authorize)
	r.NoRoute(h.NotFoundHandler)

	return r
}
