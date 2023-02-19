package conf

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type jwt struct {
	AccessTokenExpiryHour  int    `mapstructure:"access_token_expiry_hour"`
	RefreshTokenExpiryHour int    `mapstructure:"refresh_token_expiry_hour"`
	AccessTokenSecret      string `mapstructure:"access_token_secret"`
	RefreshTokenSecret     string `mapstructure:"refresh_token_secret"`
}

var JwtConf = &jwt{}

type google struct {
	ClientId     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	AuthUri      string `mapstructure:"auth_uri"`
	TokenUri     string `mapstructure:"token_uri"`
	RedirectUri  string `mapstructure:"redirect_uri"`
}

var GoogleConf = &google{}

func Setup() {
	viper.SetConfigType("YAML")
	data, err := os.ReadFile("config/config.yaml")
	if err != nil {
		logrus.Fatalf("fail to read \"config.yaml\" file: %v\n", err)
	}

	viper.ReadConfig(bytes.NewBuffer(data))

	viper.UnmarshalKey("jwt", JwtConf)
	viper.UnmarshalKey("google", GoogleConf)
}
