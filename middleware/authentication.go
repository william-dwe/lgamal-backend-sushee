package middleware

import (
	"strings"

	"final-project-backend/config"
	"final-project-backend/errorlist"
	"final-project-backend/handler/router_helper"
	"final-project-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Authenticate(c *gin.Context) {
	conf := config.Config.AuthConfig
	if conf.IsTesting != "false" {
		return
	}
	authHeader := c.GetHeader("Authorization")
	encodedTokenArray := strings.Split(authHeader, "Bearer ")
	if len(encodedTokenArray) < 2 {
		router_helper.GenerateErrorMessage(c, errorlist.UnauthorizedError())
		return
	}
	encodedToken := encodedTokenArray[1]

	a :=  utils.NewAuthUtil()
	token, err := a.ValidateToken(encodedToken, conf.HmacSecretAccessToken)
	if err != nil || !token.Valid {
		router_helper.GenerateErrorMessage(c, errorlist.UnauthorizedError())
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		router_helper.GenerateErrorMessage(c, errorlist.UnauthorizedError())
		return
	}
	c.Set("username", claims["username"])
	c.Set("role", claims["role"])
}
