package handler

import (
	"final-project-backend/entity"
	"final-project-backend/errorlist"
	"final-project-backend/handler/router_helper"
	"strconv"

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

func (h *Handler) GetOrders(c *gin.Context) {
	username := c.GetString("username")
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "0"))
	if err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("something wrong with the request content", "INVALID_INPUT"))
		return
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("something wrong with the request content", "INVALID_INPUT"))
		return
	}
	q := entity.OrderHistoryQuery{
		Search: c.DefaultQuery("s", "%"),
		SortBy: c.DefaultQuery("sortBy", "id"),
		FilterByCategory: c.DefaultQuery("filterByCategory", ""),
		Sort:   c.DefaultQuery("sort", "desc"),
		Limit:  limit,
		Page:   page,
	}


	orders, err := h.orderUsecase.GetOrderHistory(username, &q)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}

	result := entity.OrderHistoryResBody{
		Orders: *orders,
	}
	
	router_helper.GenerateResponseMessage(c, result)
}

func (h *Handler) AddReview(c *gin.Context) {
	username := c.GetString("username")
	var reqBody entity.ReviewAddReqBody
	if err := c.BindJSON(&reqBody); err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("something wrong with the request content", "INPUT_INCOMPLETE"))
		return
	}

	order, err := h.orderUsecase.AddReview(username, &reqBody)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}
	
	router_helper.GenerateResponseMessage(c, order)
}


func (h *Handler) GetOrderStatus(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "0"))
	if err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("something wrong with the request content", "INVALID_INPUT"))
		return
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("something wrong with the request content", "INVALID_INPUT"))
		return
	}
	q := entity.OrderStatusQuery{
		Search: c.DefaultQuery("s", "%"),
		SortBy: c.DefaultQuery("sortBy", "id"),
		FilterByStatus: c.DefaultQuery("filterByStatus", ""),
		Sort:   c.DefaultQuery("sort", "desc"),
		Limit:  limit,
		Page:   page,
	}


	orders, err := h.orderUsecase.GetOrderStatus(&q)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}

	result := entity.OrdersResBody{
		Orders: *orders,
	}
	
	router_helper.GenerateResponseMessage(c, result)
}


func (h *Handler) UpdateOrderStatus(c *gin.Context) {
	var reqBody entity.OrderStatusUpdateReqBody
	if err := c.BindJSON(&reqBody); err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("something wrong with the request content", "INPUT_INCOMPLETE"))
		return
	}

	orders, err := h.orderUsecase.UpdateOrderStatus(&reqBody)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}
	
	router_helper.GenerateResponseMessage(c, orders)
}