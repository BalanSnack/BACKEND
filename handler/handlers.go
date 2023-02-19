package handler

import (
	"github.com/BalanSnack/BACKEND/middleware"
	"github.com/BalanSnack/BACKEND/protocol"
	"github.com/BalanSnack/BACKEND/repo"
	"github.com/BalanSnack/BACKEND/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	state = "need to change"
)

type LoginHandler struct {
	service service.LoginService
}

type AvatarHandler struct {
	service service.AvatarService
}

func NewGinEngine() *gin.Engine {
	avatarMemStore := repo.NewAvatarMemStore()
	userMemStore := repo.NewUserMemStore()

	l := &LoginHandler{
		service: service.NewLoginService(avatarMemStore, userMemStore),
	}

	a := &AvatarHandler{
		service: service.NewAvatarService(avatarMemStore),
	}

	r := gin.Default()

	r.GET("/login/:provider", l.LoginHandler)
	r.GET("/callback/:provider", l.CallbackHandler)
	r.GET("/avatar", middleware.AuthJwt(), a.AvatarHandler)
	return r
}

func (h *LoginHandler) LoginHandler(c *gin.Context) {
	var url string

	provider := c.Params.ByName("provider")
	switch provider {
	case "google":
		url = h.service.GetGoogleLoginPageUrl(state)
	case "kakao":
	default:
	}

	c.Redirect(http.StatusTemporaryRedirect, url) // 307, 308 차이
}

func (h *LoginHandler) CallbackHandler(c *gin.Context) {
	var (
		response protocol.LoginResponse
		err      error
	)

	provider := c.Params.ByName("provider")
	switch provider {
	case "google":
		if c.Query("state") != state {
			c.String(http.StatusInternalServerError, "response's state is different with request's state")
			return
		}
		response, err = h.service.GetGoogleLoginResponse(c.Query("code"))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	case "kakao":
	default:
	}

	c.JSON(http.StatusOK, response)
}

func (h *AvatarHandler) AvatarHandler(c *gin.Context) {
	id := c.GetUint64("id")
	logrus.Traceln(id)

	avatar, err := h.service.GetAvatarById(id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, avatar)
}
