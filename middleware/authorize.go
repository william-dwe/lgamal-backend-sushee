package middleware

import (
	"encoding/json"
	"fmt"
	"strings"

	"final-project-backend/config"
	"final-project-backend/entity"
	"final-project-backend/errorlist"
	"final-project-backend/handler/router_helper"
	"final-project-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Authorize(c *gin.Context) {
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

	var userJson, _ = json.Marshal(claims["user"])
	var user entity.UserLoginReqBody
	if err := json.Unmarshal(userJson, &user); err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.UnauthorizedError())
		return
	}

	c.Set("user", user)
	fmt.Println("loggedin user: ", user)
}
