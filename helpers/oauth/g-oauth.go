package oauth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"lokasani/entity/response"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func SetupGoogleOauth() *oauth2.Config {
	oauth2GoogleConfig := &oauth2.Config{
		ClientID: viper.GetString("CLIENT_ID"),
		ClientSecret: viper.GetString("CLIENT_SECRET"),
		RedirectURL: viper.GetString("CALLBACK_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
	return oauth2GoogleConfig
}

func GetResponseAccountGoogle(code string, config *oauth2.Config) (*response.UserGoogleInfo, error) {
	token, errToken := config.Exchange(context.Background(), code)
	if errToken != nil {
		return nil, errors.New("Error exchange google")
	}

	res, errRes := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken))
	if errRes != nil {
		fmt.Println(errRes)
		return nil, errRes
	}

	defer res.Body.Close()

	var user *response.UserGoogleInfo
	err := json.NewDecoder(res.Body).Decode(&user)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	return user, nil
}