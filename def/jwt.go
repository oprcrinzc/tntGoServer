package def

import (
	"fmt"
	"os"
	"time"

	// "github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v5"
)

func GenJwt(Username string) (string, error) {
	claims := jwt.MapClaims{
		"username": Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	c := []byte(os.Getenv("jwt_secret"))
	return token.SignedString(c)
}

func VerifyJwt(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token_ *jwt.Token) (interface{}, error) {
		if _, ok := token_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		c := []byte(os.Getenv("jwt_secret"))
		return c, nil
	})
}
