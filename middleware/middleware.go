package middleware

import (
	"github.com/BalanSnack/BACKEND/conf"
	"github.com/BalanSnack/BACKEND/jwt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strings"
)

func AuthJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.Abort()
		}

		tmp := strings.Split(auth, "Bearer ")
		if len(tmp) == 2 {
			tokenString := tmp[1]
			claims, err := jwt.ParseAndValidateToken(tokenString, conf.JwtConf.AccessTokenSecret)
			logrus.Traceln(claims)
			if err != nil {
				logrus.Debug(err)
				c.Abort()
			}
			c.Set("id", uint64(claims["id"].(float64))) // https://stackoverflow.com/questions/70705673/panic-interface-conversion-interface-is-float64-not-int64
			c.Next()
		} else {
			c.Abort()
		}
	}
}
