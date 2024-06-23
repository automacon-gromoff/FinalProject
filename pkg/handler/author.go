package handler

import (
	"fmt"
	library "github.com/automacon-gromoff/FinalProject"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type getAllAuthorsResponse struct {
	Authors []library.Author `json:"authors"`
}

func (h *Handler) createAuthor(c *gin.Context) {
	var input library.Author
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	dateLayout := "2024-01-31"
	if _, err := time.Parse(dateLayout, input.BornDate); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "указана некорректная дата рождения")
		return
	}

	id, err := h.services.LibraryAuthor.Create(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllAuthors(c *gin.Context) {
	authors, err := h.services.LibraryAuthor.GetAll()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllAuthorsResponse{
		Authors: authors,
	})
}

func (h *Handler) getAuthorById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "указан некорректный id")
		return
	}

	author, err := h.services.LibraryAuthor.GetById(id)
	if err != nil {
		msg := fmt.Sprintf("автор с id = %d не найден", id)
		NewErrorResponse(c, http.StatusInternalServerError, msg)
		return
	}

	c.JSON(http.StatusOK, author)
}

func (h *Handler) updateAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "указан некорректный id")
		return
	}

	var input library.UpdateAuthorInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	dateLayout := "2024-01-31"
	if _, err := time.Parse(dateLayout, *input.BornDate); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "указана некорректная дата рождения")
		return
	}

	if err := h.services.LibraryAuthor.Update(id, input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "указан некорректный id")
		return
	}

	err = h.services.LibraryAuthor.Delete(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
