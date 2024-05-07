package auth

import (
	"strconv"
	"time"

	"github.com/edupsousa/concursos-api/config"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(userID int) (string, error) {
	secret := config.Envs.JWTSecret
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   strconv.Itoa(userID),
		"expireAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
