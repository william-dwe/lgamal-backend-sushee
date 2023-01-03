package handler

import (
	"final-project-backend/entity"
	"final-project-backend/errorlist"
	"final-project-backend/handler/router_helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddCart(c *gin.Context) {
	username := c.GetString("username")

	var reqBody entity.CartReqBody
	if err := c.BindJSON(&reqBody); err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("something wrong with the request content", "INPUT_INCOMPLETE"))
		return
	}

	cart, err := h.cartUsecase.AddCart(username, &reqBody)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}

	router_helper.GenerateResponseMessage(c, cart)
}

func (h *Handler) ShowCart(c *gin.Context) {
	username := c.GetString("username")

	t, err := h.cartUsecase.GetCart(username)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}

	resBody := entity.CartsResBody{
		Carts: *t,
	}

	router_helper.GenerateResponseMessage(c, resBody)
}

func (h *Handler) DeleteCart(c *gin.Context) {
	username := c.GetString("username")

	err := h.cartUsecase.DeleteCart(username)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}

	router_helper.GenerateResponseMessage(c, "Cart has been cleared")
}


func (h *Handler) DeleteCartById(c *gin.Context) {
	username := c.GetString("username")
	cartId, err := strconv.Atoi(c.Param("cartId"))
	if err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("wrong cartId format", "INVALID_INPUT"))
		return
	}

	err = h.cartUsecase.DeleteCartByCartId(username, cartId)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}

	router_helper.GenerateResponseMessage(c, "Cart item has been deleted")
}

func (h *Handler) UpdateCartById(c *gin.Context) {
	username := c.GetString("username")
	cartId, err := strconv.Atoi(c.Param("cartId"))
	if err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("wrong cartId format", "INVALID_INPUT"))
		return
	}
	var reqBody entity.CartEditDetailsReqBody
	if err := c.BindJSON(&reqBody); err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("should provide the new cart value", "INVALID_INPUT"))
		return
	}


	_, err = h.cartUsecase.UpdateCartByCartId(username, cartId, &reqBody)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}

	router_helper.GenerateResponseMessage(c, "cart has been updated")
}