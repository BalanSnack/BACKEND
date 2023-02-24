package app

import (
	"github.com/didnlie23/go-mvc/internals/controller"
	"github.com/didnlie23/go-mvc/internals/repository"
	"github.com/didnlie23/go-mvc/internals/service"
	"github.com/gin-gonic/gin"
)

func Run() {
	avatarMemStore := repository.NewAvatarMemStore()
	userMemStore := repository.NewUserMemStore()

	loginService := service.NewLoginService(userMemStore, avatarMemStore)
	avatarService := service.NewAvatarService(avatarMemStore)

	loginController := controller.NewLoginController(loginService)
	avatarController := controller.NewAvatarController(avatarService)

	r := gin.Default()

	r.GET("/login/:provider", loginController.Login)
	r.GET("/callback/:provider", loginController.Callback)
	r.GET("/avatar", controller.AuthJwt(), avatarController.GetByAvatarId)

	r.Run("localhost:5000")
}
