package jwt

import (
	"fmt"
	"github.com/BalanSnack/BACKEND/conf"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"time"
)

func CreateAccessToken(id uint64) (string, error) {
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * time.Duration(conf.JwtConf.AccessTokenExpiryHour)).Unix(),
	}).SignedString([]byte(conf.JwtConf.AccessTokenSecret))
	if err != nil {
		logrus.Debug(err)
		return "", err
	}

	return accessToken, nil
}

func CreateRefreshToken(id uint64) (string, error) {
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * time.Duration(conf.JwtConf.RefreshTokenExpiryHour)).Unix(),
	}).SignedString([]byte(conf.JwtConf.RefreshTokenSecret))
	if err != nil {
		logrus.Debug(err)
		return "", err
	}

	return refreshToken, nil
}

func ParseAndValidateToken(tokenString, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
