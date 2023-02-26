package util

import "github.com/gin-gonic/gin"

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	if status < 300 {
		ctx.JSON(status, er)
	} else {
		ctx.AbortWithStatusJSON(status, er)
	}
}
