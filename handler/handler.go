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
}

type HandlerConfig struct {
	UserUsecase        usecase.UserUsecase
	MenuUsecase        usecase.MenuUsecase
}

func New(c HandlerConfig) *Handler {
	return &Handler{
		userUsecase:        c.UserUsecase,
		menuUsecase: 	c.MenuUsecase,
	}
}

func (h *Handler) NotFoundHandler(c *gin.Context) {
	router_helper.GenerateErrorMessage(c, errorlist.NotFoundError("page is not found"))
}