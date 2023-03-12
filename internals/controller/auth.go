package controller

import (
	"errors"
	"github.com/BalanSnack/BACKEND/internals/entity/res"
	"github.com/BalanSnack/BACKEND/internals/service"
	"github.com/BalanSnack/BACKEND/internals/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var (
	state = "need to change" // need to fix
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// Login
// @Summary 로그인 페이지로 리다이렉트
// @Description 구글("google") 로그인 혹은 카카오("kakao") 로그인 페이지로 리다이렉트
// @Param provider path string true "provider name"
// @Success 200 {object} entity.TokenResponse
// @Failure 500 {object} util.HTTPError
// @Router /login/{provider} [get]
func (c *AuthController) Login(ctx *gin.Context) {
	var url string

	provider := ctx.Params.ByName("provider")
	switch provider {
	case "google":
		url = c.authService.GetGoogleLoginPageUrl(state)
	case "kakao":
		url = c.authService.GetKakaoLoginPageUrl(state)
	default:
	}

	ctx.Redirect(http.StatusTemporaryRedirect, url) // 307, 308 차이
}

func (c *AuthController) Callback(ctx *gin.Context) {
	var (
		response res.TokenResponse
		err      error
	)

	if ctx.Query("state") != state {
		util.NewError(ctx, http.StatusInternalServerError, errors.New("res's state is different with req's state"))
		return
	}

	provider := ctx.Params.ByName("provider")
	switch provider {
	case "google":
		response, err = c.authService.GetGoogleLoginResponse(ctx.Query("code"))
		if err != nil {
			util.NewError(ctx, http.StatusInternalServerError, err)
			return
		}
	case "kakao":
		response, err = c.authService.GetKakaoLoginResponse(ctx.Query("code"))
		if err != nil {
			util.NewError(ctx, http.StatusInternalServerError, err)
			return
		}
	default:
	}

	ctx.JSON(http.StatusOK, response)
}

// Refresh
// @Summary access token 재발급
// @Description 요청과 함께 온 리프레시 토큰이 유효한 경우, 액세스 토큰과 리프레시을 재발급
// @Security BearerAuth
// @Success 200 {object} entity.TokenResponse
// @Failure 400 {object} util.HTTPError
// @Failure 401 {object} util.HTTPError
// @Failure 406 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /refresh [get]
func (c *AuthController) Refresh(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	if header == "" {
		util.NewError(ctx, http.StatusBadRequest, errors.New("header's authorization is empty"))
		return
	}

	tmp := strings.Split(header, "Bearer ")
	if len(tmp) == 2 {
		tokenString := tmp[1]

		claims, err := util.JwtConfig.ParseAndValidateRefreshToken(tokenString)
		if err != nil {
			util.NewError(ctx, http.StatusUnauthorized, err)
			return
		}

		avatarId := uint64(claims["avatarId"].(float64))

		expiration := util.JwtConfig.CheckRefreshTokenExpiration(avatarId, tokenString)
		if !expiration {
			util.NewError(ctx, http.StatusNotAcceptable, errors.New("expired token"))
			return
		}

		accessToken, refreshToken, err := util.JwtConfig.CreateTokens(avatarId)
		if err != nil {
			util.NewError(ctx, http.StatusInternalServerError, errors.New("failed to create tokens"))
			return
		}

		ctx.JSON(http.StatusOK, res.TokenResponse{AccessToken: accessToken, RefreshToken: refreshToken})
	} else {
		util.NewError(ctx, http.StatusBadRequest, errors.New("failed to extract token string from header's authorization"))
		return
	}
}
