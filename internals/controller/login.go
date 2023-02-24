package controller

import (
	"github.com/BalanSnack/BACKEND/internals/entity"
	"github.com/BalanSnack/BACKEND/internals/service"
	"github.com/gin-gonic/gin"
	"net/http"
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
