package utility

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id int, secret string) (string, error) {
	claims := jwt.MapClaims{
		"id":  id,
		"exp": jwt.NewNumericDate(time.Now().Add(10 * time.Minute)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(token string, secret string) (jwt.MapClaims, error) {
	tokens, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(secret), nil
	})

	if err != nil {
		return jwt.MapClaims{}, err
	}

	claims, ok := tokens.Claims.(jwt.MapClaims)
	if !ok || !tokens.Valid {
		return jwt.MapClaims{}, fmt.Errorf("unable to extract claims")
	}

	return claims, nil

}
