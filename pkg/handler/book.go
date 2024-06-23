package handler

import (
	"fmt"
	library "github.com/automacon-gromoff/FinalProject"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type getAllBooksResponse struct {
	Books []library.Book `json:"books"`
}

func (h *Handler) createBook(c *gin.Context) {
	var input library.Book
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.LibraryBook.Create(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllBooks(c *gin.Context) {
	books, err := h.services.LibraryBook.GetAll()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllBooksResponse{
		Books: books,
	})
}

func (h *Handler) getBookById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "указан некорректный id")
		return
	}

	book, err := h.services.LibraryBook.GetById(id)
	if err != nil {
		msg := fmt.Sprintf("книга с id = %d не найдена", id)
		NewErrorResponse(c, http.StatusInternalServerError, msg)
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *Handler) updateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "указан некорректный id")
		return
	}

	var input library.UpdateBookInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.LibraryBook.Update(id, input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) updateBookWithAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "указан некорректный id")
		return
	}

	authorId, err := strconv.Atoi(c.Param("author_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "указан некорректный author_id")
		return
	}

	var input library.UpdateBookAndAuthorInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.LibraryBook.UpdateWithAuthor(id, authorId, input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "указан некорректный id")
		return
	}

	err = h.services.LibraryBook.Delete(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
