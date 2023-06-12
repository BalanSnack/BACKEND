package util

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

// 인-메모리에 기록한 리프레시 토큰 발급 히스토리 내역 주기적으로 비워주는 함수 필요

var JwtConfig *jwtConfig

type jwtConfig struct {
	AccessTokenExpiryHour  int
	RefreshTokenExpiryHour int
	accessTokenSecret      string
	refreshTokenSecret     string
	refreshTokenMap        map[int]map[string]bool
}

func init() {
	JwtConfig = &jwtConfig{}
}

func SetJwtUtilConfig(accessTokenExpiryHour, refreshTokenExpiryHour int, accessTokenSecret, refreshTokenSecret string) {
	JwtConfig.AccessTokenExpiryHour = accessTokenExpiryHour
	JwtConfig.RefreshTokenExpiryHour = refreshTokenExpiryHour
	JwtConfig.accessTokenSecret = accessTokenSecret
	JwtConfig.refreshTokenSecret = refreshTokenSecret
	JwtConfig.refreshTokenMap = make(map[int]map[string]bool)
}

func (ju *jwtConfig) CreateAccessToken(avatarID int, exp int64) (string, error) {
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"avatarID": avatarID,
		"exp":      exp,
	}).SignedString([]byte(ju.accessTokenSecret))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (ju *jwtConfig) CreateRefreshToken(avatarID int, exp int64) (string, error) {
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"avatarID": avatarID,
		"exp":      exp,
	}).SignedString([]byte(ju.refreshTokenSecret))
	if err != nil {
		return "", err
	}

	// 인-메모리 맵(히스토리) 불러오기
	m, ok := ju.refreshTokenMap[avatarID]
	if !ok {
		ju.refreshTokenMap[avatarID] = make(map[string]bool)
		m = ju.refreshTokenMap[avatarID]
	}

	// 기존 리프레시 토큰들 폐기
	for k := range m {
		m[k] = false
	}

	// 새로 발급할 리프레시 토큰 히스토리에 기록
	ju.refreshTokenMap[avatarID][refreshToken] = true

	return refreshToken, nil
}

func (ju *jwtConfig) ParseAndValidateAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(ju.accessTokenSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func (ju *jwtConfig) ParseAndValidateRefreshToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(ju.refreshTokenSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		log.Println(err)
		return nil, err
	}
}

func (ju *jwtConfig) CheckRefreshTokenExpiration(avatarID int, tokenString string) bool {
	return ju.refreshTokenMap[avatarID][tokenString]
}

func (ju *jwtConfig) CreateTokens(avatarID int) (string, string, error) {
	accessToken, err := ju.CreateAccessToken(avatarID, time.Now().Add(time.Hour*time.Duration(ju.AccessTokenExpiryHour)).Unix())
	if err != nil {
		return "", "", err
	}

	refreshToken, err := ju.CreateRefreshToken(avatarID, time.Now().Add(time.Hour*time.Duration(ju.RefreshTokenExpiryHour)).Unix())
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
