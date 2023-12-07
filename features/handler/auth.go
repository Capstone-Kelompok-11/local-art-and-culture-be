package handler

import (
	"net/http"
	"strings"

	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"

	"github.com/labstack/echo/v4"
)

// type GoogleOauthHandler struct {
// 	GoogleOauthService *services.GoogleOauthService
// }

// func NewGoogleOauthHandler(googleOauthService *services.GoogleOauthService) *GoogleOauthHandler {
// 	return &GoogleOauthHandler{
// 		GoogleOauthService: googleOauthService,
// 	}
// }

// func (h *GoogleOauthHandler) HandleGoogleOauthCallback(c echo.Context) error {
//     log.Println("Handling Google Oauth Callback")
//     code := c.QueryParam("code")
//     if code == "" {
//         return response.NewErrorResponse(c, echo.ErrBadGateway)
//     }

//     token, err := h.GoogleOauthService.ExchangeCode(code)
// 	if err != nil {
//     	log.Println("Error exchanging code:", err)
//     	return response.NewErrorResponse(c, echo.ErrInternalServerError)
// 	}

// 	userInfo, err := h.GoogleOauthService.GetGoogleUserInfo(token.AccessToken)
// 	if err != nil {
//     	log.Println("Error getting Google user info:", err)
//     	return response.NewErrorResponse(c, echo.ErrInternalServerError)
// 	}

//     return c.JSON(http.StatusOK, response.NewSuccessResponse(c, userInfo))
// }

type AuthHandler interface {
	AuthHandler(echo.Context) error
	OauthGoogleHandler(echo.Context) error
	OauthCallbackGoogleHandler(echo.Context) error
}

type AuthServiceImpl struct {
	AuthService services.AuthService
}

func NewAuthHandler(auth services.AuthService) AuthHandler {
	return &AuthServiceImpl{
		AuthService: auth,
	}
}

func (auth *AuthServiceImpl) AuthHandler(c echo.Context) error {
	var request request.Auth
	ErrBindRequest := c.Bind(&request)

	if ErrBindRequest != nil {
		return response.NewErrorResponse(c, ErrBindRequest)
	}
	return nil
}

func (auth *AuthServiceImpl) OauthGoogleHandler(c echo.Context) error {
	GoogleAuth := auth.AuthService.GoogleAuthService()
	return c.Redirect(http.StatusSeeOther, GoogleAuth)
}

func (auth *AuthServiceImpl) OauthCallbackGoogleHandler(c echo.Context) error {
	GetResponseEmail, errMessage := auth.AuthService.GoogleCallbackService(c)

	if errMessage != nil {
		if strings.Contains(errMessage.Error(), "state is not match") {
			return response.NewErrorResponse(c, errMessage)
		}

		if strings.Contains(errMessage.Error(), "user denied access login") {
			return response.NewErrorResponse(c, errMessage)
		}

		return response.NewErrorResponse(c, errMessage)
	}

	return response.NewSuccessResponse(c, GetResponseEmail)
}