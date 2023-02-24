package config

import (
	"github.com/didnlie23/go-mvc/internals/util"
	"github.com/spf13/viper"
	"log"
)

func Setup() {
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
}
