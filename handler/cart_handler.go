package handler

import (
	"final-project-backend/entity"
	"final-project-backend/errorlist"
	"final-project-backend/handler/router_helper"

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

	router_helper.GenerateResponseMessage(c, t)
}