package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GetSigningKey() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET environment variable not found - aborting")
	}
	return []byte(secret)
}

func GenerateJWT(id uint) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, _ := token.SignedString(GetSigningKey())
	return t
}
