package service

import (
	"context"
	"encoding/json"
	"github.com/BalanSnack/BACKEND/internals/entity"
	"github.com/BalanSnack/BACKEND/internals/repository"
	"github.com/BalanSnack/BACKEND/internals/util"
	"io"
	"log"
)

type AuthService struct {
	userRepository   repository.UserRepository
	avatarRepository repository.AvatarRepository
}

func NewAuthService(userRepository repository.UserRepository, avatarRepository repository.AvatarRepository) *AuthService {
	return &AuthService{
		userRepository:   userRepository,
		avatarRepository: avatarRepository,
	}
}

func (s *AuthService) GetGoogleLoginPageUrl(state string) string {
	return util.GoogleOAuthConfig.AuthCodeURL(state)
}

func (s *AuthService) GetGoogleLoginResponse(code string) (entity.TokenResponse, error) {
	token, err := util.GoogleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Println(err)
		return entity.TokenResponse{}, err
	}

	client := util.GoogleOAuthConfig.Client(context.Background(), token)
	res, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo") // 개선 필요
	if err != nil {
		log.Println(err)
		return entity.TokenResponse{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return entity.TokenResponse{}, err
	}

	var userInfo util.GoogleUserInfo
	if err = json.Unmarshal(data, &userInfo); err != nil {
		log.Println(err)
		return entity.TokenResponse{}, err
	}

	user, err := s.userRepository.GetByEmailAndProvider(userInfo.Email, "google")
	if err != nil {
		avatar := s.avatarRepository.Create(userInfo.Name, userInfo.Picture, false)
		user = s.userRepository.Create(avatar.AvatarId, userInfo.Email, "google")
	}

	accessToken, err := util.JwtConfig.CreateAccessToken(user.AvatarId)
	if err != nil {
		log.Println(err)
		return entity.TokenResponse{}, err
	}
	refreshToken, err := util.JwtConfig.CreateRefreshToken(user.AvatarId)
	if err != nil {
		log.Println(err)
		return entity.TokenResponse{}, err
	}

	return entity.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) getGoogleUserInfo(code string) (util.GoogleUserInfo, error) {
	token, err := util.GoogleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Println(err)
		return util.GoogleUserInfo{}, err
	}

	client := util.GoogleOAuthConfig.Client(context.Background(), token)
	res, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo") // 개선 필요
	if err != nil {
		log.Println(err)
		return util.GoogleUserInfo{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return util.GoogleUserInfo{}, err
	}

	var userInfo util.GoogleUserInfo
	if err = json.Unmarshal(data, &userInfo); err != nil {
		log.Println(err)
		return util.GoogleUserInfo{}, err
	}

	return userInfo, nil
}
