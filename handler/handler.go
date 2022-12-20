package handler

import (
	"final-project-backend/errorlist"
	"final-project-backend/handler/router_helper"
	"final-project-backend/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userUsecase        usecase.UserUsecase
}

type HandlerConfig struct {
	UserUsecase        usecase.UserUsecase
}

func New(c HandlerConfig) *Handler {
	return &Handler{
		userUsecase:        c.UserUsecase,
	}
}

func (h *Handler) NotFoundHandler(c *gin.Context) {
	router_helper.GenerateErrorMessage(c, errorlist.NotFoundError("page is not found"))
}