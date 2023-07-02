package controller

import (
	"BACKEND/internals/entity"
	"BACKEND/internals/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GameController 게임 생성 및 조회
type GameController struct {
	gameService *service.GameService
}

func NewGameController(gameService *service.GameService) *GameController {
	return &GameController{
		gameService: gameService,
	}
}

func (c *GameController) Get(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	avatarID := ctx.GetInt("avatarID")

	// 첫 화면에서 API 호출 시 id == 0
	game, err := c.gameService.Get(id, avatarID, ctx.Param("class"))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, game)
}

func (c *GameController) Create(ctx *gin.Context) {
	req := entity.CreateGameRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	avatarID := ctx.GetInt("avatarID")

	game, err := c.gameService.Create(avatarID, req)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, game)
}

func (c *GameController) Update(ctx *gin.Context) {
	req := entity.UpdateGameRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	game, err := c.gameService.Update(req)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	game.AvatarID = ctx.GetInt("avatarID")

	ctx.JSON(http.StatusOK, game)
}

func (c *GameController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = c.gameService.Delete(id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err) // 에러 정의해서 반환 코드 분기 필요
		return
	}

	ctx.Status(http.StatusOK)
}
