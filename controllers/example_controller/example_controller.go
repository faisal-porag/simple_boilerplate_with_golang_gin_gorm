package example_controller

import (
	"github.com/gin-gonic/gin"
	"golang_boilerplate_with_gin/models"
	"net/http"
)

type AddBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h *handler) AddBook(ctx *gin.Context) {
	body := AddBookRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	if result := h.DB.Create(&book); result.Error != nil {
		_ = ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &book)
}

func (h *handler) GetBooks(ctx *gin.Context) {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		_ = ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &books)
}

func (h *handler) GetBook(ctx *gin.Context) {
	id := ctx.Param("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		_ = ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &book)
}

type UpdateBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h *handler) UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	body := UpdateBookRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		_ = ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	h.DB.Save(&book)

	ctx.JSON(http.StatusOK, &book)
}

func (h *handler) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		_ = ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&book)

	ctx.Status(http.StatusOK)
}
