package controller

import (
	"errors"
	"github.com/BalanSnack/BACKEND/internals/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func CheckAccessToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("Authorization")
		if auth == "" {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("information required for authentication is missing; Authorization header"))
			return
		}

		tmp := strings.Split(auth, "Bearer ")
		if len(tmp) == 2 {
			tokenString := tmp[1]
			claims, err := util.JwtConfig.ParseAndValidateAccessToken(tokenString)
			if err != nil {
				ctx.AbortWithError(http.StatusUnauthorized, err)
				return
			}
			ctx.Set("avatarID", int(claims["avatarID"].(float64)))
			ctx.Next()
		} else {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("invalid value of Authorization header"))
			return
		}
	}
}
