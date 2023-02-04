package auth

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig *oauth2.Config = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     "524616912419-9anfecerga0bfa03n0b0fqososkc4hmh.apps.googleusercontent.com",
		ClientSecret: "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",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	googleOauthStateString string = "didnlie23's random string"
)

func GetGoogleOauthLogin(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL(googleOauthStateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func GetGoogleOauthCallback(c *gin.Context) {
	state := c.Query("state")
	if state != googleOauthStateString {
		return
	}
	code := c.Query("code")
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, token)
}
