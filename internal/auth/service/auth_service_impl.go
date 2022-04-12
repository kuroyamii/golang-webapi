package authServicePkg

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type authService struct {
}

func ProvideAuthService() authService {
	return authService{}
}

var AUTH_KEY = []byte("auth-key")

func (as *authService) CreateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute).Unix()

	tokenString, err := token.SignedString(AUTH_KEY)
	if err != nil {
		log.Println(err.Error())
		return "Error", nil
	}
	return tokenString, nil
}
