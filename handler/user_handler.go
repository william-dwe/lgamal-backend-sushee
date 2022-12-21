package handler

import (
	"final-project-backend/config"
	"final-project-backend/errorlist"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"final-project-backend/entity"
	"final-project-backend/handler/router_helper"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

func (h *Handler) UserDetail(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	s := strings.Split(authHeader, "Bearer ")
	if len(s) < 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	encodedToken := s[1]
	
	var username string
	token, _, err := new(jwt.Parser).ParseUnverified(encodedToken, jwt.MapClaims{})
	if err != nil {
		fmt.Printf("Error %s", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		username= fmt.Sprint(claims["username"])
	}

	userInfo, err := h.userUsecase.GetDetailUserByUsername(username)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}

	router_helper.GenerateResponseMessage(c, userInfo)
}


func (h *Handler) UpdateUser(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	s := strings.Split(authHeader, "Bearer ")
	if len(s) < 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	encodedToken := s[1]
	
	var username string
	token, _, err := new(jwt.Parser).ParseUnverified(encodedToken, jwt.MapClaims{})
	if err != nil {
		fmt.Printf("Error %s", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		username= fmt.Sprint(claims["username"])
	}


	var reqBody entity.UserEditDetailsReqBody
	if err := c.BindJSON(&reqBody); err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("should provide the changes", "UPDATE_INPUT_INCOMPLETE"))
		return
	}

	userInfo, err := h.userUsecase.UpdateUserDetailsByUsername(username, reqBody)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}

	router_helper.GenerateResponseMessage(c, userInfo)
}