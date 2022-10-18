package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var SecretKey = []byte("af6afsgsd")

type MyClaim struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func NewToken(name string) (string, error) {
	user := MyClaim{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
			Issuer:    "Yogen",
		},
		name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	return token.SignedString(SecretKey)
}

func AuthToken(tokenString string) (*MyClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaim{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, is_ok := token.Claims.(*MyClaim)
	if !token.Valid || !is_ok {
		return nil, fmt.Errorf("token不合法")
	}
	return claims, nil
}
