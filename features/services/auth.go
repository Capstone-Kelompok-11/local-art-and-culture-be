package services

import (
	"errors"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	"lokasani/entity/domain"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/generate"
	"lokasani/helpers/oauth"
	"lokasani/helpers/bcrypt"
	"lokasani/helpers/middleware"
)

// type IAuthGoogleService interface {
// 	ExchangeCode(code string) (*oauth2.Token, error)
// 	GetGoogleUserInfo(accessToken string) (*response.UserGoogleInfo, error)
// }

// type GoogleOauthService struct {
// 	Config *oauth2.Config
// }

// func NewAuthGoogleService(Config *oauth2.Config) *GoogleOauthService {
//     return &GoogleOauthService{Config: Config}
// }

// func (g *GoogleOauthService) ExchangeCode(code string) (*oauth2.Token, error) {
//     log.Println("Exchanging code:", code)
//     return g.Config.Exchange(context.Background(), code)
// }

// func (g *GoogleOauthService) GetGoogleUserInfo(accessToken string) (*response.UserGoogleInfo, error) {
//     log.Println("Getting user info with access token:", accessToken)
// 	res, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(accessToken))
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, err
// 	}
// 	defer res.Body.Close()

// 	var user *response.UserGoogleInfo
// 	err = json.NewDecoder(res.Body).Decode(&user)
// 	if err != nil {
// 		fmt.Print(err)
// 		return nil, err
// 	}
// 	return user, nil
// }

type IAuthService interface {
	GoogleAuthService() string
	GoogleCallbackService(echo.Context) (*response.AuthToken, error)
}

type AuthService struct {
	RoleRepo repositories.IRoleRepository
	UserRepo repositories.IUserRepository
	validate *validator.Validate
}

func NewAuthService(role repositories.IRoleRepository, user repositories.IUserRepository, validate *validator.Validate) *AuthService{
	return &AuthService{
		RoleRepo: role,
		UserRepo: user,
		validate: validate,
	}
}

func (auth *AuthService) GoogleAuthService() string {
	googleOauth := oauth.SetupGoogleOauth()
	url := googleOauth.AuthCodeURL(viper.GetString("STATE_STRING"))
	return url
}

func (auth *AuthService) GoogleCallbackService(ctx echo.Context) (*response.AuthToken, error) {
	var SetAuthenticateData response.Auth
	var GetAuthWithTokenResponse *response.AuthToken

	StateQuery := ctx.FormValue("state")
	CodeQuery := ctx.FormValue("code")

	if StateQuery != viper.GetString("STATE_STRING") {
		return nil, errors.New("state does not match")
	}

	if CodeQuery == "" {
		return nil, errors.New("user denied access login")
	}

	googleOauthService := oauth.SetupGoogleOauth()

	userGoogleInfo, err := oauth.GetResponseAccountGoogle(CodeQuery, googleOauthService)
	if err != nil {
		return nil, errors.New("error getting Google user info")
	}

	UserResponse, ErrExists := auth.UserRepo.FindByEmail(userGoogleInfo.Email)

	if ErrExists != nil {
		SetUser := request.User{
			FirstName:      UserResponse.FirstName,
			LastName:       UserResponse.LastName,
			Email:           UserResponse.Email,
			Password:        helpers.StringWithCharset(10), 
			PhoneNumber:    "62123456789",
		}

		GetRole, _ := auth.RoleRepo.FindByName("users")

		if GetRole == nil {
			return nil, errors.New("role not found")
		}

		SetUser.RoleId = uint(GetRole.ID)

		UserConvert := domain.UserCreateRequestToUserDomain(SetUser)
		hashedPassword, _ := bcrypt.HashPassword(SetUser.Password)

		UserConvert.Password = hashedPassword

		UserCreate, ErrCreate := auth.UserRepo.CreateUser(UserConvert)

		if ErrCreate != nil {
			return nil, errors.New("failed when processing user data")
		}

		GetUserByEmail, ErrGetUser := auth.UserRepo.FindByEmail(UserCreate.Email)

		if ErrGetUser != nil {
			return nil, errors.New("failed when processing user data")
		}

		SetAuthenticateData = domain.ConvertUserToAuthRes(GetUserByEmail)

	} else {
		SetAuthenticateData = domain.ConvertUserToAuthRes(UserResponse)
	}

	userID := SetAuthenticateData.Id
	GetTokenAuth, ErrGetToken := middleware.CreateToken(userID, 0, 0, "", "")

	if ErrGetToken != nil {
		return nil, ErrGetToken
	}

	GetAuthWithTokenResponse = domain.AuthResToAuthTokenRes(SetAuthenticateData, GetTokenAuth)

	return GetAuthWithTokenResponse, nil
}