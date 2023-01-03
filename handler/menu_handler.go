package handler

import (
	"final-project-backend/entity"
	"final-project-backend/errorlist"
	"final-project-backend/handler/router_helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ShowMenu(c *gin.Context) {
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

	respBody := entity.PromotionResBody{
		Promotions: *t,
	}

	router_helper.GenerateResponseMessage(c, respBody)
}


func (h *Handler) AddMenu(c *gin.Context) {
	var reqBody entity.MenuAddReqBody
	if err := c.BindJSON(&reqBody); err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("something wrong with the request content", "INPUT_INCOMPLETE"))
		return
	}

	
	menu, err := h.menuUsecase.AddMenu(&reqBody)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}

	router_helper.GenerateResponseMessage(c, menu)
}


func (h *Handler) UpdateMenu(c *gin.Context) {
	menuId, err := strconv.Atoi(c.Param("menuId"))
	if err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("wrong menuId format", "INVALID_INPUT"))
		return
	}
	var reqBody entity.MenuAddReqBody
	if err := c.BindJSON(&reqBody); err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("should provide the new menu value", "INVALID_INPUT"))
		return
	}

	newMenu := entity.Menu{
		MenuName: reqBody.MenuName,
		Price: reqBody.Price,
		MenuPhoto: reqBody.MenuPhoto,
		CategoryId: reqBody.CategoryId,
		Customization: reqBody.Customization,
	}

	menu, err := h.menuUsecase.UpdateMenuByMenuId(menuId, &newMenu)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}

	router_helper.GenerateResponseMessage(c, menu)
}

func (h *Handler) DeleteMenu(c *gin.Context) {
	menuId, err := strconv.Atoi(c.Param("menuId"))
	if err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("wrong menuId format", "INVALID_INPUT"))
		return
	}

	err = h.menuUsecase.DeleteMenuByMenuId(menuId)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}

	router_helper.GenerateResponseMessage(c, "Menu item has been deleted")
}

func (h *Handler) GetMenuDetail(c *gin.Context) {
	menuId, err := strconv.Atoi(c.Param("menuId"))
	if err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("wrong menuId format", "INVALID_INPUT"))
		return
	}

	menu, err := h.menuUsecase.GetMenuDetailByMenuId(menuId)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}

	router_helper.GenerateResponseMessage(c, menu)
}