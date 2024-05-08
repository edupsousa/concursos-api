package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/edupsousa/concursos-api/config"
	user_model "github.com/edupsousa/concursos-api/services/user/model"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserKey contextKey = "userID"

func CreateJWT(userID uint) (string, error) {
	secret := config.Envs.JWTSecret
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   strconv.FormatUint(uint64(userID), 10),
		"expireAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func WithJWTAuth(handlerFunc http.HandlerFunc, store user_model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := getTokenFromRequest(r)
		token, err := validateTokenString(tokenString)
		if err != nil {
			log.Printf("error validating token: %v", err)
			permissionDenied(w)
			return
		}
		if !token.Valid {
			log.Printf("invalid token")
			permissionDenied(w)
			return
		}
		user, err := getUserFromToken(token, store)
		if err != nil {
			log.Printf("failed to get user: %v", err)
			permissionDenied(w)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, user.ID)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}

func getUserFromToken(token *jwt.Token, store user_model.UserStore) (*user_model.User, error) {
	claims := token.Claims.(jwt.MapClaims)
	strUserID := claims["userID"].(string)
	userID, _ := strconv.Atoi(strUserID)
	user := store.GetUserByID(userID)
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func getTokenFromRequest(r *http.Request) string {
	return r.Header.Get("Authorization")
}

func validateTokenString(tokenString string) (*jwt.Token, error) {
	secret := config.Envs.JWTSecret
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	http.Error(w, "permission denied", http.StatusForbidden)
}

func GetUserIDFromContext(ctx context.Context) int {
	userID := ctx.Value(UserKey)
	if userID == nil {
		return 0
	}
	return userID.(int)
}
