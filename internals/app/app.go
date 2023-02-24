package app

import (
	"github.com/BalanSnack/BACKEND/internals/controller"
	"github.com/BalanSnack/BACKEND/internals/repository"
	"github.com/BalanSnack/BACKEND/internals/service"
	"github.com/gin-gonic/gin"
)

func Run() {
	avatarMemStore := repository.NewAvatarMemStore()
	userMemStore := repository.NewUserMemStore()

	authService := service.NewAuthService(userMemStore, avatarMemStore)
	avatarService := service.NewAvatarService(avatarMemStore)

	authController := controller.NewAuthController(authService)
	avatarController := controller.NewAvatarController(avatarService)

	r := gin.Default()

	r.GET("/login/:provider", authController.Login)
	r.GET("/callback/:provider", authController.Callback)
	r.GET("/avatar", controller.CheckAccessToken(), avatarController.GetByAvatarId)
	r.GET("/refresh")

	r.Run("localhost:5000")
}
