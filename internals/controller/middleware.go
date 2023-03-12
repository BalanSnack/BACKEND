package controller

import (
	"github.com/BalanSnack/BACKEND/internals/util"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func CheckAccessToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("Authorization")
		if auth == "" {
			ctx.Abort()
		}

		tmp := strings.Split(auth, "Bearer ")
		if len(tmp) == 2 {
			tokenString := tmp[1]
			claims, err := util.JwtConfig.ParseAndValidateAccessToken(tokenString)
			if err != nil {
				log.Println(err)
				ctx.Abort()
			}
			ctx.Set("avatarId", uint64(claims["avatarId"].(float64)))
			ctx.Next()
		} else {
			ctx.Abort()
		}
	}
}
