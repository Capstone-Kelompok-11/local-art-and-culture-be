package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func CreateToken(userId uint, role string, creatorId uint) (string, error) {
	godotenv.Load()
	claims := jwt.MapClaims{}
	claims["id"] = userId
	claims["role_id"] = role
	claims["creator_id"] = creatorId
	// claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["role"] = role

	fmt.Println("Token Claims:", claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}

func JWTMiddleware() echo.MiddlewareFunc {
	godotenv.Load()
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(os.Getenv("SECRET_JWT")),
		SigningMethod: "HS256",
		//TokenLookup:   "cookie:token",
	})
}

func SetTokenCookie(e echo.Context, token string) {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Path = "/"

	e.SetCookie(cookie)
}

func ExtractToken(e echo.Context) (uint, string, uint, error) {
	user, ok := e.Get("user").(*jwt.Token)
	if !ok {
		return 0, "", 0, errors.New("invalid token")
	}

	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return 0, "", 0, errors.New("invalid token claims")
	}

	// exp, ok := claims["exp"].(float64)
	// if !ok || time.Now().Unix() > int64(exp) {
	// 	return 0, 0, 0, "", "", errors.New("token has expired")
	// }

	userIDFloat, ok := claims["id"].(float64)
	if !ok {
		return 0, "", 0, errors.New("invalid token claims")
	}
	userID := uint(userIDFloat)

	role, ok := claims["role_id"].(string)
	if !ok {
		return 0, "", 0, errors.New("invalid token claims")
	}
	roleID := role

	creatorIDFloat, ok := claims["creator_id"].(float64)
	if !ok {
		return 0, "", 0, errors.New("invalid token claims")
	}
	creatorID := uint(creatorIDFloat)

	// name, okName := claims["name"].(string)
	// if !okName {
	// 	return 0, 0, 0, "", "", errors.New("invalid token claims")
	// }

	// role, okRole := claims["role"].(string)
	// if !okRole {
	// 	return 0, 0, 0, "", errors.New("invalid token claims")
	// }

	fmt.Println("Token Claims:", claims)
	return userID, roleID, creatorID, nil
}
