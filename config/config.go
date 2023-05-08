package config

import (
	"github.com/spf13/viper"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/kakao"
	"log"
)

type Config struct {
	GoogleOAuthConfig GoogleOAuthConfig
	KakaoOAuthConfig  KakaoOAuthConfig
	JwtConfig         JwtConfig
}

func NewConfig() (*Config, error) {
	viper.SetConfigFile("config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	newGoogleConfig := GoogleOAuthConfig{
		ClientId:     viper.GetString("google.client_id"),
		ClientSecret: viper.GetString("google.client_secret"),
		Endpoint:     google.Endpoint,
		RedirectUri:  viper.GetString("google.redirect_uri"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	}

	newKakaoConfig := KakaoOAuthConfig{
		ClientId:     viper.GetString("kakao.client_id"),
		ClientSecret: viper.GetString("kakao.client_secret"),
		Endpoint:     kakao.Endpoint,
		RedirectUri:  viper.GetString("google.redirect_uri"),
	}

	newJWTConfig := JwtConfig{
		AccessTokenExpiryHour:  viper.GetInt("jwt.access_token_expiry_hour"),
		RefreshTokenExpiryHour: viper.GetInt("jwt.refresh_token_expiry_hour"),
		accessTokenSecret:      viper.GetString("jwt.access_token_secret"),
		refreshTokenSecret:     viper.GetString("jwt.refresh_token_secret"),
	}

	return &Config{
		GoogleOAuthConfig: newGoogleConfig,
		KakaoOAuthConfig:  newKakaoConfig,
		JwtConfig:         newJWTConfig,
	}, nil
}
