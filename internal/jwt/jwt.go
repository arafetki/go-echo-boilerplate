package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func HMACCheck(tokenString string, secret string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing algorithm: %s", t.Method.Alg())
		}
		return []byte(secret), nil
	})
}

func RSACheck(tokenString string, publicKey string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing algorithm: %s", t.Method.Alg())
		}
		return []byte(publicKey), nil
	})
}
