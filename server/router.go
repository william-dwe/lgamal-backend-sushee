package server

import (
	"final-project-backend/handler"
	"final-project-backend/middleware"
	"final-project-backend/usecase"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	UserUsecase        usecase.UserUsecase
	MenuUsecase 	usecase.MenuUsecase
	CartUsecase usecase.CartUsecase
	OrderUsecase usecase.OrderUsecase
	CouponUsecase usecase.CouponUsecase
}

func CreateRouter(c RouterConfig) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.JSONifyResult())

	h := handler.New(handler.HandlerConfig{
		UserUsecase:        c.UserUsecase,
		MenuUsecase: c.MenuUsecase,
		CartUsecase: c.CartUsecase,
		OrderUsecase: c.OrderUsecase,
		CouponUsecase: c.CouponUsecase,
	})

	v1 := r.Group("/v1")
	v1.GET("/menus", h.ShowMenu)
	v1.GET("/promotions", h.ShowPromotion)
	v1.POST("/login", h.Login)
	v1.POST("/register", h.Register)
	v1.GET("/refresh", h.Refresh)

	user := v1.Group("")
	user.Use(middleware.Authorize)
	user.GET("/users/me", h.ShowUserDetail)
	user.POST("/users/me", h.UpdateUserProfile)

	user.GET("/carts", h.ShowCart)
	user.GET("/carts/:cartId", h.ShowCartById)
	user.POST("/carts", h.AddCart)
	user.POST("/carts/:cartId", h.UpdateCartById)
	user.DELETE("/carts", h.DeleteCart)
	user.DELETE("/carts/:cartId", h.DeleteCartById)

	user.GET("/orders", h.GetOrder)
	user.POST("/orders", h.AddOrder)
	user.GET("/orders/payment", h.GetPaymentOption)
	user.GET("/orders/coupon", h.GetUserCouponByUsername)

	admin := v1.Group("")
	admin.Use(middleware.Authorize)
	admin.POST("/coupons", h.AddCoupon)
	admin.GET("/coupons", h.GetCoupon)
	admin.POST("/coupons/:couponId", h.UpdateCoupon)
	admin.DELETE("/coupons/:couponId", h.DeleteCoupon)


	r.NoRoute(h.NotFoundHandler)
	return r
}
