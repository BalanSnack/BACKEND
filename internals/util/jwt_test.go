package util

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	// setup
	SetJwtUtilConfig(1, 1, "access_token_secret", "refresh_token_secret")

	test := m.Run()

	os.Exit(test)
}

func TestJwtConfig_CreateAccessToken(t *testing.T) {
	avatarID := 1024
	exp := time.Now().Add(time.Hour * time.Duration(JwtConfig.AccessTokenExpiryHour)).Unix() // int to time.Duration

	// 액세스 토큰 생성
	token, err := JwtConfig.CreateAccessToken(avatarID, exp)
	if err != nil {
		t.Fatal(err)
	}

	// 파싱 + 유효성 검사
	result, err := JwtConfig.ParseAndValidateAccessToken(token)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, avatarID, int(result["avatarID"].(float64)))
	assert.Equal(t, exp, int64(result["exp"].(float64)))
}

// https://stackoverflow.com/questions/70705673/panic-interface-conversion-interface-is-float64-not-int64

func TestJwtConfig_CreateRefreshToken(t *testing.T) {
	avatarID := 1024
	exp := time.Now().Add(time.Hour * time.Duration(JwtConfig.RefreshTokenExpiryHour)).Unix()

	// 리프레시 토큰 생성
	token, err := JwtConfig.CreateRefreshToken(avatarID, exp)
	if err != nil {
		t.Fatal(err)
	}

	// 유효성 검사
	result, err := JwtConfig.ParseAndValidateRefreshToken(token)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, avatarID, int(result["avatarID"].(float64)))
	assert.Equal(t, exp, int64(result["exp"].(float64)))
}

func TestJwtConfig_CheckRefreshTokenExpiration(t *testing.T) {
	avatarID := 1024
	exp := time.Now().Add(time.Hour * time.Duration(JwtConfig.RefreshTokenExpiryHour)).Unix()

	oldToken, err := JwtConfig.CreateRefreshToken(avatarID, exp)
	if err != nil {
		t.Fatal(err)
	}

	newExp := time.Now().Add(time.Hour*time.Duration(JwtConfig.RefreshTokenExpiryHour) + time.Second).Unix() // 초 단위로 값에 차이가 발생

	newToken, err := JwtConfig.CreateRefreshToken(avatarID, newExp)
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, JwtConfig.CheckRefreshTokenExpiration(avatarID, newToken))
	assert.False(t, JwtConfig.CheckRefreshTokenExpiration(avatarID, oldToken))
}
