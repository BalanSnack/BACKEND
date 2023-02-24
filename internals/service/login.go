package service

import (
	"context"
	"encoding/json"
	"github.com/didnlie23/go-mvc/internals/entity"
	"github.com/didnlie23/go-mvc/internals/repository"
	"github.com/didnlie23/go-mvc/internals/util"
	"io"
	"log"
)

type LoginService struct {
	userRepository   repository.UserRepository
	avatarRepository repository.AvatarRepository
}

func NewLoginService(userRepository repository.UserRepository, avatarRepository repository.AvatarRepository) *LoginService {
	return &LoginService{
		userRepository:   userRepository,
		avatarRepository: avatarRepository,
	}
}

func (s *LoginService) GetGoogleLoginPageUrl(state string) string {
	return util.GoogleOAuthConfig.AuthCodeURL(state)
}

func (s *LoginService) GetGoogleLoginResponse(code string) (entity.LoginResponse, error) {
	token, err := util.GoogleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Println(err)
		return entity.LoginResponse{}, err
	}

	client := util.GoogleOAuthConfig.Client(context.Background(), token)
	res, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo") // 개선 필요
	if err != nil {
		log.Println(err)
		return entity.LoginResponse{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return entity.LoginResponse{}, err
	}

	var userInfo util.GoogleUserInfo
	if err = json.Unmarshal(data, &userInfo); err != nil {
		log.Println(err)
		return entity.LoginResponse{}, err
	}

	user, err := s.userRepository.GetByEmailAndProvider(userInfo.Email, "google")
	if err != nil {
		avatar := s.avatarRepository.Create(userInfo.Name, userInfo.Picture, false)
		user = s.userRepository.Create(avatar.AvatarId, userInfo.Email, "google")
	}

	accessToken, err := util.JwtConfig.CreateAccessToken(user.AvatarId)
	if err != nil {
		log.Println(err)
		return entity.LoginResponse{}, err
	}
	refreshToken, err := util.JwtConfig.CreateRefreshToken(user.AvatarId)
	if err != nil {
		log.Println(err)
		return entity.LoginResponse{}, err
	}

	return entity.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *LoginService) getGoogleUserInfo(code string) (util.GoogleUserInfo, error) {
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
