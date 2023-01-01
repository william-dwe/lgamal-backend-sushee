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
	orderUsecase usecase.OrderUsecase
	couponUsecase usecase.CouponUsecase
}

type HandlerConfig struct {
	UserUsecase        usecase.UserUsecase
	MenuUsecase        usecase.MenuUsecase
	CartUsecase usecase.CartUsecase
	OrderUsecase usecase.OrderUsecase
	CouponUsecase usecase.CouponUsecase
}

func New(c HandlerConfig) *Handler {
	return &Handler{
		userUsecase:        c.UserUsecase,
		menuUsecase: 	c.MenuUsecase,
		cartUsecase: c.CartUsecase,
		orderUsecase: c.OrderUsecase,
		couponUsecase: c.CouponUsecase,
	}
}

func (h *Handler) NotFoundHandler(c *gin.Context) {
	router_helper.GenerateErrorMessage(c, errorlist.NotFoundError("page is not found"))
}