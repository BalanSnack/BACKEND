package config

import (
	"log"
	"os"

	docs "BACKEND/docs"
	"BACKEND/internals/util"

	"github.com/spf13/viper"
)

const (
	EnvLocal = "local"
	EnvDev   = "dev"
)

func Setup() {
	env := os.Getenv("ENV")
	if env == "" {
		env = EnvLocal
	}

	switch env {
	case EnvLocal:
		viper.SetConfigFile("config.yaml")

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal(err)
		}

		util.SetJwtUtilConfig(
			viper.GetInt("jwt.access_token_expiry_hour"),
			viper.GetInt("jwt.refresh_token_expiry_hour"),
			viper.GetString("jwt.access_token_secret"),
			viper.GetString("jwt.refresh_token_secret"))

		util.SetGoogleOAuthConfig(
			viper.GetString("google.client_id"),
			viper.GetString("google.client_secret"),
			viper.GetString("google.redirect_uri"))

		util.SetKakaoOAuthConfig(
			viper.GetString("kakao.client_id"),
			viper.GetString("kakao.client_secret"),
			viper.GetString("kakao.redirect_uri"))

		docs.SwaggerInfo.Title = "BalanSnack Server API"
	case EnvDev:
	default:
		log.Fatal("invalid environment")
	}
}
