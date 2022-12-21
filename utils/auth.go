package utils

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"final-project-backend/config"
)

type AuthUtil interface {
	GenerateRefreshToken() (string, error)
	GenerateAccessToken(username string) (string, error)
	ValidateToken(encodedToken, signSecret string) (*jwt.Token, error)
	ComparePassword(hashedPwd string, inputPwd string) bool
}

type authUtilImpl struct{}

func NewAuthUtil() AuthUtil {
	return &authUtilImpl{}
}

var c = config.Config.AuthConfig
type customAccessTokenClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (a *authUtilImpl) GenerateAccessToken(username string) (string, error) {
	expirationLimit, _ := strconv.ParseInt(c.TimeLimitAccessToken, 10, 64)
	claims := &customAccessTokenClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expirationLimit))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(c.HmacSecretAccessToken))

	return tokenString, err
}

type customRefreshTokenClaims struct {
	jwt.RegisteredClaims
}

func (a *authUtilImpl) GenerateRefreshToken() (string, error) {
	expirationLimit, _ := strconv.ParseInt(c.TimeLimitRefreshToken, 10, 64)
	claims := &customRefreshTokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expirationLimit))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(c.HmacSecretRefreshToken))

	return tokenString, err
}

func (a *authUtilImpl) ValidateToken(encodedToken, signSecret string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, errors.New("invalid signature token")
		}
		return []byte(signSecret), nil
	})	
}


func (a *authUtilImpl) ComparePassword(hashedPwd string, inputPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(inputPwd))
	return err == nil
}
