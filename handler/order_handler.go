package handler

import (
	"final-project-backend/entity"
	"final-project-backend/errorlist"
	"final-project-backend/handler/router_helper"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPaymentOption(c *gin.Context) {
	paymentOptions, err := h.orderUsecase.GetPaymentOption()
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}

	resBody := entity.PaymentOptionResBody{
		PaymentOptions: *paymentOptions,
	}

	router_helper.GenerateResponseMessage(c, &resBody)
}

func (h *Handler) AddOrder(c *gin.Context) {
	username := c.GetString("username")
	var reqBody entity.OrderReqBody
	if err := c.BindJSON(&reqBody); err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("something wrong with the request content", "INPUT_INCOMPLETE"))
		return
	}

	order, err := h.orderUsecase.AddOrder(username, &reqBody)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}
	
	router_helper.GenerateResponseMessage(c, order)
}

func (h *Handler) DeleteOrder(c *gin.Context) {
	username := c.GetString("username")
	var reqBody entity.OrderReqBody
	if err := c.BindJSON(&reqBody); err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("something wrong with the request content", "INPUT_INCOMPLETE"))
		return
	}

	order, err := h.orderUsecase.AddOrder(username, &reqBody)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}
	
	router_helper.GenerateResponseMessage(c, order)
}

func (h *Handler) GetOrder(c *gin.Context) {
	username := c.GetString("username")
	orders, err := h.orderUsecase.GetOrderHistory(username)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}
	
	router_helper.GenerateResponseMessage(c, orders)
}

