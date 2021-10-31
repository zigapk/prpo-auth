package util

import (
	"crypto/rsa"
	"errors"

	"github.com/dgrijalva/jwt-go"
)

var ErrInvalidToken = errors.New("invalid token")

type Claims struct {
	jwt.StandardClaims
}

var signingKey *rsa.PublicKey

func Validate(accessToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(_ *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, ErrInvalidToken
	}
}

func SetKey(key *rsa.PublicKey) {
	signingKey = key
}
