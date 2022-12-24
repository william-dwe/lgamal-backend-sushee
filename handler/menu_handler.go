package handler

import (
	"final-project-backend/entity"
	"final-project-backend/errorlist"
	"final-project-backend/handler/router_helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ShowMenu(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "5"))
	if err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("something wrong with the request content", "INVALID_INPUT"))
		return
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("something wrong with the request content", "INVALID_INPUT"))
		return
	}

	q := entity.MenuQuery{
		Search: c.DefaultQuery("s", "%"),
		SortBy: c.DefaultQuery("sortBy", "category_id"),
		FilterByCategory: c.DefaultQuery("filterByCategory", ""),
		Sort:   c.DefaultQuery("sort", "desc"),
		Limit:  limit,
		Page:   page,
	}

	t, maxPage, err := h.menuUsecase.GetMenu(q)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}

	respBody := entity.MenuResBody{
		Menus: *t,
		CurrentPage: q.Page,
		MaxPage: maxPage,
	}

	router_helper.GenerateResponseMessage(c, respBody)
}

func (h *Handler) ShowPromotion(c *gin.Context) {
	t, err := h.menuUsecase.GetPromotion()
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}

	router_helper.GenerateResponseMessage(c, t)
}

// todo add menu detail