package util

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOAuthConfig *oauth2.Config

func init() {
	GoogleOAuthConfig = &oauth2.Config{
		Scopes:   []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}
}

type GoogleUserInfo struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
}

func SetGoogleOAuthConfig(clientId, clientSecret, redirectUri string) {
	GoogleOAuthConfig.ClientID = clientId
	GoogleOAuthConfig.ClientSecret = clientSecret
	GoogleOAuthConfig.RedirectURL = redirectUri
}
