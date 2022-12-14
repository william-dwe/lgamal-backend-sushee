package utils

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"final-project-backend/config"
	"final-project-backend/entity"
)

type AuthUtil interface {
	GenerateAccessToken(user *entity.User) (string, error)
	ComparePassword(hashedPwd string, inputPwd string) bool
}

type authUtilImpl struct{}

func NewAuthUtil() AuthUtil {
	return &authUtilImpl{}
}

type customClaims struct {
	User *entity.User `json:"user"`
	jwt.RegisteredClaims
}

var c = config.Config.AuthConfig

func (a *authUtilImpl) GenerateAccessToken(user *entity.User) (string, error) {
	expirationLimit, _ := strconv.ParseInt(c.ExpiresAt, 10, 64)
	claims := &customClaims{
		user,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expirationLimit))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	hmacSampleSecret := c.HmacSecret
	tokenString, err := token.SignedString([]byte(hmacSampleSecret))

	return tokenString, err
}

func (a *authUtilImpl) ComparePassword(hashedPwd string, inputPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(inputPwd))
	return err == nil
}
