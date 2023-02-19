package service

import (
	"context"
	"encoding/json"
	"github.com/BalanSnack/BACKEND/conf"
	"github.com/BalanSnack/BACKEND/jwt"
	"github.com/BalanSnack/BACKEND/protocol"
	"github.com/BalanSnack/BACKEND/repo"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io"
)

type LoginService struct {
	googleOAuthConfig *oauth2.Config
	avatarRepo        repo.AvatarRepo
	userRepo          repo.UserRepo
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

func NewLoginService(avatarRepo repo.AvatarRepo, userRepo repo.UserRepo) LoginService {
	config := &oauth2.Config{
		RedirectURL:  conf.GoogleConf.RedirectUri,
		ClientID:     conf.GoogleConf.ClientId,
		ClientSecret: conf.GoogleConf.ClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	return LoginService{
		googleOAuthConfig: config,
		avatarRepo:        avatarRepo,
		userRepo:          userRepo,
	}
}

func (s *LoginService) GetGoogleLoginPageUrl(state string) string {
	return s.googleOAuthConfig.AuthCodeURL(state)
}

func (s *LoginService) GetGoogleLoginResponse(code string) (protocol.LoginResponse, error) {
	userInfo, err := s.getGoogleUserInfo(code)
	if err != nil {
		logrus.Debug(err)
		return protocol.LoginResponse{}, err
	}

	var avatar repo.Avatar

	user, err := s.userRepo.GetUserByEmailAndProvider(userInfo.Email, "google")
	if err != nil {
		avatar = s.avatarRepo.Create(userInfo.Name, userInfo.Picture, false)
		user = s.userRepo.Create(avatar.Id, userInfo.Email, "google")
	} else {
		avatar, err = s.avatarRepo.GetAvatarById(user.AvatarId)
		if err != nil {
			logrus.Debug(err)
			return protocol.LoginResponse{}, err
		}
	}

	accessToken, err := jwt.CreateAccessToken(user.AvatarId)
	if err != nil {
		logrus.Debug(err)
		return protocol.LoginResponse{}, err
	}
	refreshToken, err := jwt.CreateRefreshToken(user.AvatarId)
	if err != nil {
		logrus.Debug(err)
		return protocol.LoginResponse{}, err
	}

	return protocol.LoginResponse{
		Id:           avatar.Id,
		Nickname:     avatar.Nickname,
		Profile:      avatar.Profile,
		Anonymity:    avatar.Anonymity,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *LoginService) getGoogleUserInfo(code string) (GoogleUserInfo, error) {
	token, err := s.googleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		logrus.Debug(err)
		return GoogleUserInfo{}, err
	}

	client := s.googleOAuthConfig.Client(context.Background(), token)
	res, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo") // 개선 필요
	if err != nil {
		logrus.Debug(err)
		return GoogleUserInfo{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		logrus.Debug(err)
		return GoogleUserInfo{}, err
	}

	var userInfo GoogleUserInfo
	if err = json.Unmarshal(data, &userInfo); err != nil {
		logrus.Debug(err)
		return GoogleUserInfo{}, err
	}

	return userInfo, nil
}
