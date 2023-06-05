package controller

import (
	"github.com/BalanSnack/BACKEND/internals/repository"
	"github.com/BalanSnack/BACKEND/internals/repository/mysql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type CommentController struct {
	commentRepository *mysql.CommentRepository
}

func NewCommentController(commentRepository *mysql.CommentRepository) *CommentController {
	return &CommentController{
		commentRepository: commentRepository,
	}
}

func (c *CommentController) Create(ctx *gin.Context) {
	comment := repository.Comment{}
	err := ctx.ShouldBindJSON(&comment)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = c.commentRepository.Create(&comment)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, comment)
}

func (c *CommentController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = c.commentRepository.Delete(id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *CommentController) Update(ctx *gin.Context) {
	req := struct {
		ID      int    `json:"id"`
		Content string `json:"content"`
	}{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	log.Println(req)

	err = c.commentRepository.Update(req.ID, req.Content)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}
