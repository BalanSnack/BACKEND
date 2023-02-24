package controller

import (
	"github.com/BalanSnack/BACKEND/internals/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AvatarController struct {
	avatarService *service.AvatarService
}

func NewAvatarController(avatarService *service.AvatarService) *AvatarController {
	return &AvatarController{
		avatarService: avatarService,
	}
}

func (c *AvatarController) GetByAvatarId(ctx *gin.Context) {
	avatarId := ctx.GetUint64("avatarId")

	avatar, err := c.avatarService.GetByAvatarId(avatarId)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, avatar)
}
