package config

import (
	"golang.org/x/oauth2"
)

type GoogleOAuthConfig struct {
	ClientId     string
	ClientSecret string
	Endpoint     oauth2.Endpoint
	RedirectUri  string
	Scopes       []string
	MyInfoUri    string
}

type KakaoOAuthConfig struct {
	ClientId     string
	ClientSecret string
	Endpoint     oauth2.Endpoint
	RedirectUri  string
	Scopes       []string
	MyInfoUri    string
}
