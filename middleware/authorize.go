package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"final-project-backend/config"
	"final-project-backend/entity"
	"final-project-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)


func Authorize(c *gin.Context) {
	conf := config.Config.AuthConfig
	if conf.IsTesting != "false" {
		fmt.Println("disabled JWT auth for testing")
		return
	}
	authHeader := c.GetHeader("Authorization")
	s := strings.Split(authHeader, "Bearer ")
	if len(s) < 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	decodedtoken := s[1]

	a :=  utils.NewAuthUtil()
	token, err := a.ValidateToken(decodedtoken, conf.HmacSecretAccessToken)
	if err != nil || !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var userJson, _ = json.Marshal(claims["user"])
	var user entity.UserLoginReqBody
	if err := json.Unmarshal(userJson, &user); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("user", user)
	fmt.Println("loggedin user: ", user)
}
