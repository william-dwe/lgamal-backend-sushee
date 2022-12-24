package handler

import (
	"final-project-backend/errorlist"
	"final-project-backend/handler/router_helper"
	"final-project-backend/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userUsecase        usecase.UserUsecase
	menuUsecase 		usecase.MenuUsecase
	cartUsecase	usecase.CartUsecase
}

type HandlerConfig struct {
	UserUsecase        usecase.UserUsecase
	MenuUsecase        usecase.MenuUsecase
	CartUsecase usecase.CartUsecase
}

func New(c HandlerConfig) *Handler {
	return &Handler{
		userUsecase:        c.UserUsecase,
		menuUsecase: 	c.MenuUsecase,
		cartUsecase: c.CartUsecase,
	}
}

func (h *Handler) NotFoundHandler(c *gin.Context) {
	router_helper.GenerateErrorMessage(c, errorlist.NotFoundError("page is not found"))
}