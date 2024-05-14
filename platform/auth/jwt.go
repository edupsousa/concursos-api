package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/edupsousa/concursos-api/platform/config"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserKey contextKey = "userID"
const RoleKey contextKey = "roleID"

type JWTUser interface {
	GetID() uint
	GetRoleID() uint
}

type JWTUserRepository interface {
	FindByID(id uint) JWTUser
}

func CreateJWT(user JWTUser) (string, error) {
	secret := config.Envs.JWTSecret
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   strconv.FormatUint(uint64(user.GetID()), 10),
		"expireAt": time.Now().Add(expiration).Unix(),
		"roleID":   strconv.FormatUint(uint64(user.GetRoleID()), 10),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func WithJWTAuth(handlerFunc http.HandlerFunc, repo JWTUserRepository) http.HandlerFunc {
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
		user, err := getUserFromToken(token, repo)
		if err != nil {
			log.Printf("failed to get user: %v", err)
			permissionDenied(w)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, user.GetID())
		ctx = context.WithValue(ctx, RoleKey, user.GetRoleID())
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}

func getUserFromToken(token *jwt.Token, userRepo JWTUserRepository) (JWTUser, error) {
	claims := token.Claims.(jwt.MapClaims)
	strUserID := claims["userID"].(string)
	userID, _ := strconv.ParseUint(strUserID, 10, 64)
	user := userRepo.FindByID(uint(userID))
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
