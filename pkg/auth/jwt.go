package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtClaims struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func CreateJWTClaims(id string, name string) *jwtClaims {
	return &jwtClaims{
		id,
		name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
		},
	}
}

func GenerateToken(claims jwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("SECRET_KEY")

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}
