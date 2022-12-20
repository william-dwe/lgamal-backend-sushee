package handler

import (
	"final-project-backend/config"
	"final-project-backend/errorlist"
	"strconv"

	"final-project-backend/entity"
	"final-project-backend/handler/router_helper"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {
	var reqBody entity.UserRegisterReqBody

	if err := c.BindJSON(&reqBody); err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("something wrong with the request content", "INPUT_INCOMPLETE"))
		return
	}

	u, err := h.userUsecase.Register(&reqBody)

	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}

	router_helper.GenerateResponseMessage(c, u)
}

func (h *Handler) Login(c *gin.Context) {
	var reqBody entity.UserLoginReqBody
	if err := c.BindJSON(&reqBody); err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("should provide identifier and password", "LOGIN_INPUT_INCOMPLETE"))
		return
	}

	AccessToken, refreshToken, err := h.userUsecase.Login(reqBody.Identifier, reqBody.Password)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}

	expirationLimit, _ := strconv.Atoi(config.Config.AuthConfig.TimeLimitRefreshToken)
	c.SetCookie(
		"refreshToken",
		refreshToken,
		expirationLimit,
		"/",
		"localhost",
		false,
		true,
	)

	router_helper.GenerateResponseMessage(c, AccessToken)
}


func (h *Handler) Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}

	accessToken, err := h.userUsecase.Refresh(refreshToken)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}
	router_helper.GenerateResponseMessage(c, accessToken)
}

// func (h *Handler) ShowUserDetail(c *gin.Context) {
// 	user, ok := c.Get("user")
// 	if !ok {
// 		router_helper.GenerateErrorMessage(c, errorlist.UnauthorizedError())
// 		return
// 	}
// 	userId := user.(entity.UserContext).Id

// 	userInfo, err := h.userUsecase.GetDetailUser(userId)
// 	if err != nil {
// 		router_helper.GenerateErrorMessage(c, err)
// 		return
// 	}

// 	router_helper.GenerateResponseMessage(c, userInfo)
// }