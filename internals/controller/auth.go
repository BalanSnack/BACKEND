package controller

import (
	"errors"
	"github.com/BalanSnack/BACKEND/internals/entity"
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

func (c *AuthController) Login(ctx *gin.Context) {
	var url string

	provider := ctx.Params.ByName("provider")
	switch provider {
	case "google":
		url = c.authService.GetGoogleLoginPageUrl(state)
	case "kakao":
	default:
	}

	ctx.Redirect(http.StatusTemporaryRedirect, url) // 307, 308 차이
}

func (c *AuthController) Callback(ctx *gin.Context) {
	var (
		response entity.TokenResponse
		err      error
	)

	provider := ctx.Params.ByName("provider")
	switch provider {
	case "google":
		if ctx.Query("state") != state {
			ctx.String(http.StatusInternalServerError, "response's state is different with request's state")
			return
		}
		response, err = c.authService.GetGoogleLoginResponse(ctx.Query("code"))
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
	case "kakao":
	default:
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *AuthController) Refresh(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	if header == "" {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("header's authorization is empty"))
		return
	}

	tmp := strings.Split(header, "Bearer ")
	if len(tmp) == 2 {
		tokenString := tmp[1]

		claims, err := util.JwtConfig.ParseAndValidateRefreshToken(tokenString)
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		avatarId := uint64(claims["avatarId"].(float64))

		expiration := util.JwtConfig.CheckTokenExpiration(avatarId, tokenString)
		if !expiration {
			ctx.AbortWithError(http.StatusNotAcceptable, errors.New("expired token"))
			return
		}

		accessToken, err := util.JwtConfig.CreateAccessToken(avatarId)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		refreshToken, err := util.JwtConfig.CreateRefreshToken(avatarId)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, entity.TokenResponse{AccessToken: accessToken, RefreshToken: refreshToken})
	} else {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("failed to extract token string from header's authorization"))
		return
	}
}
