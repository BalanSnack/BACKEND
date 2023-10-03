package app

import (
	"database/sql"

	"BACKEND/internals/controller"
	"BACKEND/internals/pkg/mysql"
	"BACKEND/internals/service"

	"github.com/gin-gonic/gin"
)

func Run() {
	db, err := sql.Open("mysql", "balansnack:balansnack@tcp(localhost:3306)/balansnack?parseTime=true")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	memberRepository := mysql.NewMemberRepository(db)
	avatarRepository := mysql.NewAvatarRepository(db)
	gameRepository := mysql.NewGameRepository(db)
	voteRepository := mysql.NewVoteRepository(db)
	commentRepository := mysql.NewCommentRepository(db)
	likeRepository := mysql.NewLikeRepository(db)

	authService := service.NewAuthService(memberRepository, avatarRepository)
	avatarService := service.NewAvatarService(avatarRepository)
	gameService := service.NewGameService(gameRepository, likeRepository, voteRepository, commentRepository)

	authController := controller.NewAuthController(authService)
	avatarController := controller.NewAvatarController(avatarService)
	gameController := controller.NewGameController(gameService)

	commentController := controller.NewCommentController(commentRepository)
	likeController := controller.NewLikeController(likeRepository)
	voteController := controller.NewVoteController(voteRepository)

	r := gin.Default()

	mw := controller.CheckAccessToken()

	r.GET("/login/:provider", authController.Login)
	r.GET("/callback/:provider", authController.Callback)
	r.GET("/refresh", authController.Refresh)

	r.GET("/avatar", mw, avatarController.GetByAvatarId)

	r.GET("/game/:class/:id", mw, gameController.Get)
	r.POST("/game", mw, gameController.Create)
	r.DELETE("/game/:id", mw, gameController.Delete)
	r.PUT("/game", mw, gameController.Update)

	r.POST("/comment", mw, commentController.Create)
	r.DELETE("/comment/:id", mw, commentController.Delete)
	r.PUT("/comment", mw, commentController.Update)

	r.POST("/like", mw, likeController.Create)
	r.DELETE("/like/:class/:id", mw, likeController.Delete) // 수정 필요

	r.POST("/vote", mw, voteController.Create)
	r.DELETE("/vote/:id", mw, voteController.Delete) // 사용 X 예정

	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
