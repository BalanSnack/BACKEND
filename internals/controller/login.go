package controller

import (
	"github.com/didnlie23/go-mvc/internals/entity"
	"github.com/didnlie23/go-mvc/internals/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	state = "need to change" // need to fix
)

type LoginController struct {
	loginService *service.LoginService
}

func NewLoginController(loginService *service.LoginService) *LoginController {
	return &LoginController{
		loginService: loginService,
	}
}

func (c *LoginController) Login(ctx *gin.Context) {
	var url string

	provider := ctx.Params.ByName("provider")
	switch provider {
	case "google":
		url = c.loginService.GetGoogleLoginPageUrl(state)
	case "kakao":
	default:
	}

	ctx.Redirect(http.StatusTemporaryRedirect, url) // 307, 308 차이
}

func (c *LoginController) Callback(ctx *gin.Context) {
	var (
		response entity.LoginResponse
		err      error
	)

	provider := ctx.Params.ByName("provider")
	switch provider {
	case "google":
		if ctx.Query("state") != state {
			ctx.String(http.StatusInternalServerError, "response's state is different with request's state")
			return
		}
		response, err = c.loginService.GetGoogleLoginResponse(ctx.Query("code"))
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
	case "kakao":
	default:
	}

	ctx.JSON(http.StatusOK, response)
}
