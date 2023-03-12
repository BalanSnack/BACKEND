package controller

import (
	"github.com/BalanSnack/BACKEND/internals/entity/req"
	"github.com/BalanSnack/BACKEND/internals/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GameController struct {
	gameService *service.GameService
}

func NewGameController(gameService *service.GameService) *GameController {
	return &GameController{
		gameService: gameService,
	}
}

func (c *GameController) GetAll(ctx *gin.Context) {
	games := c.gameService.GetAll()

	ctx.JSON(http.StatusOK, games)
}

func (c *GameController) GetByTagId(ctx *gin.Context) {
	tagId, err := strconv.ParseUint(ctx.Param("tag-id"), 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	games := c.gameService.GetByTagId(tagId)

	ctx.JSON(http.StatusOK, games)
}

func (c *GameController) GetByGameId(ctx *gin.Context) {
	gameId, err := strconv.ParseUint(ctx.Param("game-id"), 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	game, err := c.gameService.GetByGameId(gameId)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, game)
}

func (c *GameController) Create(ctx *gin.Context) {
	req := req.CreateGameRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	avatarId := ctx.GetUint64("avatarId")

	game, err := c.gameService.Create(avatarId, req)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, game)
}

func (c *GameController) Update(ctx *gin.Context) {
	req := req.UpdateGameRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	avatarId := ctx.GetUint64("avatarId")

	game, err := c.gameService.Update(avatarId, req)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err) // 에러 정의해서 반환 코드 분기 필요
		return
	}

	ctx.JSON(http.StatusOK, game)
}

func (c *GameController) Delete(ctx *gin.Context) {
	gameId, err := strconv.ParseUint(ctx.Param("game-id"), 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	avatarId := ctx.GetUint64("avatarId")

	err = c.gameService.Delete(avatarId, gameId)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err) // 에러 정의해서 반환 코드 분기 필요
		return
	}

	ctx.Status(http.StatusOK)
}
