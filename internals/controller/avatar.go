package controller

import (
	"github.com/BalanSnack/BACKEND/internals/service"
	"github.com/BalanSnack/BACKEND/internals/util"
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

// @BasePath /avatar

// GetByAvatarId
// @Summary 아바타 정보 조회
// @Description 액세스 토큰에서 추출한 `avatarId`를 활용해 아바타 정보 조회
// @Security BearerAuth
// @Success 200 {object} repository.Avatar
// @Failure 500 {object} util.HTTPError
// @Router /avatar [get]
func (c *AvatarController) GetByAvatarId(ctx *gin.Context) {
	avatarId := ctx.GetUint64("avatarId")

	avatar, err := c.avatarService.GetByAvatarId(avatarId)
	if err != nil {
		util.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, avatar)
}
