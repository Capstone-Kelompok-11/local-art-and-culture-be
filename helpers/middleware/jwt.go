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

func CreateToken(userId uint, name, role string) (string, error) {
    godotenv.Load()
    claims := jwt.MapClaims{}
    claims["id"] = userId
    claims["name"] = name
    claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
    claims["role"] = role
    fmt.Println(claims)
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

func ExtractToken(e echo.Context) (uint, string, string, error) {
    user := e.Get("user").(*jwt.Token)
    if user.Valid {
        claims := user.Claims.(jwt.MapClaims)

        userIDFloat, ok := claims["id"].(float64)
        if !ok {
            return 0, "", "", errors.New("invalid token claims")
        }

        userID := uint(userIDFloat)
        name, okName := claims["name"].(string)
        if !okName {
            return 0, "", "", errors.New("invalid token claims")
        }

        role, okRole := claims["role"].(string)
        if !okRole {
            return 0, "", "", errors.New("invalid token claims")
        }

        e.Set("userID", userID)
        e.Set("userName", name)
        e.Set("userRole", role)

        return userID, name, role, nil
    }
    return 0, "", "", errors.New("invalid token")
}