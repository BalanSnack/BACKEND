package service

import (
	"BACKEND/internals/entity"
	"BACKEND/internals/pkg"
	"BACKEND/internals/pkg/mysql"
	"BACKEND/internals/util"
	"context"
	"encoding/json"
	"io"
	"log"
)

type AuthService struct {
	memberRepository *mysql.MemberRepository
	avatarRepository *mysql.AvatarRepository
}

func NewAuthService(memberRepository *mysql.MemberRepository, avatarRepository *mysql.AvatarRepository) *AuthService {
	return &AuthService{
		memberRepository: memberRepository,
		avatarRepository: avatarRepository,
	}
}

func (s *AuthService) GetKakaoLoginPageUrl(state string) string {
	return util.KakaoOAuthConfig.AuthCodeURL(state)
}

func (s *AuthService) GetKakaoLoginResponse(code string) (entity.TokenResponse, error) {
	token, err := util.KakaoOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Println(err)
		return entity.TokenResponse{}, err
	}

	client := util.KakaoOAuthConfig.Client(context.Background(), token)
	response, err := client.Get("https://kapi.kakao.com/v2/user/me")
	if err != nil {
		log.Println(err)
		return entity.TokenResponse{}, err
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			// what should i do?
		}
	}()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return entity.TokenResponse{}, err
	}

	var userInfo util.KakaoUser
	if err = json.Unmarshal(data, &userInfo); err != nil {
		log.Println(err)
		return entity.TokenResponse{}, err
	}

	member, err := s.memberRepository.GetByEmailAndProvider(userInfo.KakaoAccount.Email, "kakao")
	if err != nil {
		log.Println(err)
		return entity.TokenResponse{}, err
	}
	if member == nil {
		avatar := &pkg.Avatar{
			Nick:    userInfo.KakaoAccount.Profile.Nickname,
			Profile: userInfo.KakaoAccount.Profile.ProfileImageURL,
		}
		err = s.avatarRepository.Create(avatar)
		if err != nil {
			log.Println(err)
			return entity.TokenResponse{}, err
		}
		member = &pkg.Member{
			Email:    userInfo.KakaoAccount.Email,
			Provider: "kakao",
			AvatarID: avatar.ID,
		}
		err = s.memberRepository.Create(member)
		if err != nil {
			log.Println(err)
			return entity.TokenResponse{}, err
		}
	}

	accessToken, refreshToken, err := util.JwtConfig.CreateTokens(member.AvatarID)
	if err != nil {
		return entity.TokenResponse{}, err
	}

	return entity.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
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
	response, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo") // 개선 필요
	if err != nil {
		log.Println(err)
		return entity.TokenResponse{}, err
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			// what should i do?
		}
	}()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return entity.TokenResponse{}, err
	}

	var userInfo util.GoogleUserInfo
	if err = json.Unmarshal(data, &userInfo); err != nil {
		log.Println(err)
		return entity.TokenResponse{}, err
	}

	member, err := s.memberRepository.GetByEmailAndProvider(userInfo.Email, "google")
	if err != nil {
		log.Println(err)
		return entity.TokenResponse{}, err
	}
	if member == nil {
		avatar := &pkg.Avatar{
			Nick:    userInfo.Name,
			Profile: userInfo.Picture,
		}
		err = s.avatarRepository.Create(avatar)
		if err != nil {
			log.Println(err)
			return entity.TokenResponse{}, err
		}
		member = &pkg.Member{
			Email:    userInfo.Email,
			Provider: "google",
			AvatarID: avatar.ID,
		}
		err = s.memberRepository.Create(member)
		if err != nil {
			log.Println(err)
			return entity.TokenResponse{}, err
		}
	}

	log.Println(member.AvatarID)
	accessToken, refreshToken, err := util.JwtConfig.CreateTokens(member.AvatarID)
	if err != nil {
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
	response, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo") // 개선 필요
	if err != nil {
		log.Println(err)
		return util.GoogleUserInfo{}, err
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			// what should i do?
		}
	}()

	data, err := io.ReadAll(response.Body)
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
