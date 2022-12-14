package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"final-project-backend/config"
	"final-project-backend/entity"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func validateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, errors.New("invalid signature token")
		}
		return []byte("very-secret"), nil
	})
}

func Authorize(c *gin.Context) {
	if config.Config.AuthConfig.IsTesting != "false" {
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

	token, err := validateToken(decodedtoken)
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
