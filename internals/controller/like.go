package controller

import (
	"fmt"
	"github.com/BalanSnack/BACKEND/internals/repository"
	"github.com/BalanSnack/BACKEND/internals/repository/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// LikeController 생성, 삭제
type LikeController struct {
	likeRepository *mysql.LikeRepository
}

func NewLikeController(likeRepository *mysql.LikeRepository) *LikeController {
	return &LikeController{
		likeRepository: likeRepository,
	}
}

func (c *LikeController) Create(ctx *gin.Context) {
	like := repository.Like{}
	err := ctx.ShouldBindJSON(&like)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	like.AvatarID = ctx.GetInt("avatarID")

	if like.GameID != 0 && like.CommentID != 0 {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid request; only one of game_id or comment_id is required"))
		return
	} else if like.GameID != 0 {
		err = c.likeRepository.CreateLikeGame(&like)
	} else if like.CommentID != 0 {
		err = c.likeRepository.CreateLikeComment(&like)
	} else {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid request; one of game_id or comment_id is required"))
		return
	}
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, like)
}

func (c *LikeController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	avatarID := ctx.GetInt("avatarID")

	switch ctx.Param("class") {
	case "game":
		err = c.likeRepository.DeleteByGameID(id, avatarID)
	case "comment":
		err = c.likeRepository.DeleteByCommentID(id, avatarID)
	default:
		err = fmt.Errorf("invalid class %v; 'comment' or 'game' are required as the value of class", ctx.Param("class"))
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}
