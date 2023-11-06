package helpers

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token 	:= c.Get("user").(*jwt.Token)
			claims 	:= token.Claims.(jwt.MapClaims)
			role 	:= claims["role"].(string)

			_, ok := claims["role"].(string)
			if !ok {
   				 return c.JSON(http.StatusForbidden, "Role not found in token claims")
			}

			for _, allowedRole := range roles {
    			if role == allowedRole {
        			return next(c)
    			}
			}

			return c.JSON(http.StatusForbidden, "Unauthorized")
		}
	}
}