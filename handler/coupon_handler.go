package handler

import (
	"final-project-backend/entity"
	"final-project-backend/errorlist"
	"final-project-backend/handler/router_helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddCoupon(c *gin.Context) {
	username := c.GetString("username")

	var reqBody entity.CouponAddReqBody
	if err := c.BindJSON(&reqBody); err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("something wrong with the request content", "INPUT_INCOMPLETE"))
		return
	}

	coupon, err := h.couponUsecase.AddCoupon(username, &reqBody)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}
	router_helper.GenerateResponseMessage(c, coupon)
}


func (h *Handler) GetCoupon(c *gin.Context) {
	coupons, err := h.couponUsecase.GetCoupon()
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}
	router_helper.GenerateResponseMessage(c, coupons)
}

func (h *Handler) UpdateCoupon(c *gin.Context) {
	username := c.GetString("username")
	couponId, err := strconv.Atoi(c.Param("couponId"))
	if err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("wrong cartId format", "INVALID_INPUT"))
		return
	}

	var reqBody entity.CouponEditReqBody
	if err := c.BindJSON(&reqBody); err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("something wrong with the request content", "INPUT_INCOMPLETE"))
		return
	}
	
	coupon, err := h.couponUsecase.UpdateCoupon(username, couponId, &reqBody)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}

	router_helper.GenerateResponseMessage(c, coupon)
}

func (h *Handler) DeleteCoupon(c *gin.Context) {
	couponId, err := strconv.Atoi(c.Param("couponId"))
	if err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("wrong cartId format", "INVALID_INPUT"))
		return
	}
	coupons, err := h.couponUsecase.DeleteCoupon(couponId)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}
	router_helper.GenerateResponseMessage(c, coupons)
}


func (h *Handler) AddUserCoupon(c *gin.Context) {
	var reqBody entity.UserCouponAddReqBody
	if err := c.BindJSON(&reqBody); err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("something wrong with the request content", "INPUT_INCOMPLETE"))
		return
	}

	

	// todo: input userid & couponid; body json?
	// logic: check coupon avial -> create usercoupon -> deduct coupon slot

	// what's next?
	// 1. test to integrate coupon with order
	// 2. create coupon component at profile & cart
	// 3. trx history + review
	// 4. favorited menu
	// CRUD TIME ADMIN!!!!
	// UT

	// coupons, err := h.couponUsecase.GetUserCouponByUsername(username)
	// if err != nil {
	// 	router_helper.GenerateErrorMessage(c, err)
	// 	return
	// }
	// router_helper.GenerateResponseMessage(c, coupons)
}



func (h *Handler) GetUserCouponByUsername(c *gin.Context) {
	username := c.GetString("username")

	coupons, err := h.couponUsecase.GetUserCouponByUsername(username)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}
	router_helper.GenerateResponseMessage(c, coupons)
}

