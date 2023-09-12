package example_controller2

import (
	"github.com/gin-gonic/gin"
	"golang_boilerplate_with_gin/models"
	"golang_boilerplate_with_gin/state"
	"net/http"
)

type AddBookRequestBodyExample struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func HandleAddBookExample(s *state.State, ctx *gin.Context) {
	body := AddBookRequestBodyExample{}

	if err := ctx.BindJSON(&body); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	if result := s.Repository.CreateBook(&book); result.Error != nil {
		_ = ctx.AbortWithError(http.StatusNotFound, result)
		return
	}

	ctx.JSON(http.StatusCreated, &book)
}
