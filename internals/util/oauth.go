package util

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/kakao"
	"time"
)

var (
	GoogleOAuthConfig *oauth2.Config
	KakaoOAuthConfig  *oauth2.Config
)

func init() {
	GoogleOAuthConfig = &oauth2.Config{
		Scopes:   []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}
	KakaoOAuthConfig = &oauth2.Config{
		Endpoint: kakao.Endpoint,
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

type KakaoAccount struct {
	ProfileNicknameNeedsAgreement bool `json:"profile_nickname_needs_agreement"`
	ProfileImageNeedsAgreement    bool `json:"profile_image_needs_agreement"`
	Profile                       struct {
		Nickname          string `json:"nickname"`
		ThumbnailImageURL string `json:"thumbnail_image_url"`
		ProfileImageURL   string `json:"profile_image_url"`
		IsDefaultImage    bool   `json:"is_default_image"`
	} `json:"profile"`
	HasEmail            bool   `json:"has_email"`
	EmailNeedsAgreement bool   `json:"email_needs_agreement"`
	IsEmailValid        bool   `json:"is_email_valid"`
	IsEmailVerified     bool   `json:"is_email_verified"`
	Email               string `json:"email"`
}

type KakaoUser struct {
	ID          int64     `json:"id"`
	ConnectedAt time.Time `json:"connected_at"`
	Properties  struct {
		Nickname       string `json:"nickname"`
		ProfileImage   string `json:"profile_image"`
		ThumbnailImage string `json:"thumbnail_image"`
	} `json:"properties"`
	KakaoAccount KakaoAccount `json:"kakao_account"`
}

func SetKakaoOAuthConfig(clientId, clientSecret, redirectUri string) {
	KakaoOAuthConfig.ClientID = clientId
	KakaoOAuthConfig.ClientSecret = clientSecret
	KakaoOAuthConfig.RedirectURL = redirectUri
}
