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
	r.Use(middleware.JSONifyResult())

	h := handler.New(handler.HandlerConfig{
		UserUsecase:        c.UserUsecase,
	})

	v1 := r.Group("/v1")
	v1.POST("/login", h.Login)
	v1.POST("/register", h.Register)
	v1.GET("/refresh", h.Refresh)

	user := v1.Group("")
	user.Use(middleware.Authorize)
	user.GET("/users/me", h.UserDetail)
	user.POST("/users/me", h.UpdateUser)
	user.POST("/users/me/profile", h.UpdateUserProfile)
	
	r.NoRoute(h.NotFoundHandler)

	return r
}
