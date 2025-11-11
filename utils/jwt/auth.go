package jwt

import (
	"time"

	"github.com/ahsansaif47/cdc-app/config"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.GetConfig().JWTSecret)

type Claims struct {
	Email    string
	UserName string
	Role     uint
	UserID   string
	jwt.RegisteredClaims
}

func GenerateJWT(userID, email, userName string, roleID uint) (string, error) {
	expTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Email:    email,
		UserName: userName,
		Role:     roleID,
		UserID:   userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "home-kitchens",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
