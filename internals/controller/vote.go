package controller

import (
	"BACKEND/internals/pkg"
	"BACKEND/internals/pkg/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// VoteController 게임 참여 생성, 삭제
type VoteController struct {
	voteRepository *mysql.VoteRepository
}

func NewVoteController(voteRepository *mysql.VoteRepository) *VoteController {
	return &VoteController{
		voteRepository: voteRepository,
	}
}

func (c *VoteController) Create(ctx *gin.Context) {
	vote := pkg.Vote{}
	err := ctx.ShouldBindJSON(&vote)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	vote.AvatarID = ctx.GetInt("avatarID")

	err = c.voteRepository.Create(&vote)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, vote)
}

func (c *VoteController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = c.voteRepository.Delete(id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}
