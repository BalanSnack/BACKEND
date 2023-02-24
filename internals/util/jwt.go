package util

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

var JwtConfig *jwtConfig

type jwtConfig struct {
	accessTokenExpiryHour  int
	refreshTokenExpiryHour int
	accessTokenSecret      string
	refreshTokenSecret     string
}

func init() {
	JwtConfig = &jwtConfig{}
}

func SetJwtUtilConfig(accessTokenExpiryHour, refreshTokenExpiryHour int, accessTokenSecret, refreshTokenSecret string) {
	JwtConfig.accessTokenExpiryHour = accessTokenExpiryHour
	JwtConfig.refreshTokenExpiryHour = refreshTokenExpiryHour
	JwtConfig.accessTokenSecret = accessTokenSecret
	JwtConfig.refreshTokenSecret = refreshTokenSecret
}

func (ju *jwtConfig) CreateAccessToken(avatarId uint64) (string, error) {
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"avatarId": avatarId,
		"exp":      time.Now().Add(time.Hour * time.Duration(ju.accessTokenExpiryHour)).Unix(),
	}).SignedString([]byte(ju.accessTokenSecret))
	if err != nil {
		log.Println(err)
		return "", err
	}

	return accessToken, nil
}

func (ju *jwtConfig) CreateRefreshToken(avatarId uint64) (string, error) {
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"avatarId": avatarId,
		"exp":      time.Now().Add(time.Hour * time.Duration(ju.refreshTokenExpiryHour)).Unix(),
	}).SignedString([]byte(ju.refreshTokenSecret))
	if err != nil {
		log.Println(err)
		return "", err
	}

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
		log.Println(err)
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
