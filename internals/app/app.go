package app

import (
	_ "github.com/BalanSnack/BACKEND/docs"
	"github.com/BalanSnack/BACKEND/internals/controller"
	"github.com/BalanSnack/BACKEND/internals/repository"
	"github.com/BalanSnack/BACKEND/internals/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run() {
	avatarMemStore := repository.NewAvatarMemStore()
	userMemStore := repository.NewUserMemStore()
	gameMemStore := repository.NewGameMemStore()

	authService := service.NewAuthService(userMemStore, avatarMemStore)
	avatarService := service.NewAvatarService(avatarMemStore)
	gameService := service.NewGameService(gameMemStore)

	authController := controller.NewAuthController(authService)
	avatarController := controller.NewAvatarController(avatarService)
	gameController := controller.NewGameController(gameService)

	r := gin.Default()

	r.GET("/login/:provider", authController.Login)
	r.GET("/callback/:provider", authController.Callback)
	r.GET("/refresh", authController.Refresh)

	r.GET("/avatar", controller.CheckAccessToken(), avatarController.GetByAvatarId)

	r.GET("/game", gameController.GetAll)
	r.GET("/game/:game-id", gameController.GetByGameId)
	r.GET("/game/tag/:tag-id", gameController.GetByTagId)
	r.POST("/game", controller.CheckAccessToken(), gameController.Create)
	r.DELETE("/game/:game-id", controller.CheckAccessToken(), gameController.Delete)
	r.PUT("/game", controller.CheckAccessToken(), gameController.Update)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run("localhost:5000")
}
