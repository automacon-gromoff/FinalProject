package handler

import (
	"github.com/automacon-gromoff/FinalProject/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		books := api.Group("/books")
		{
			books.POST("/", h.createBook)
			books.GET("/", h.getAllBooks)
			books.GET("/:id", h.getBookById)
			books.PUT("/:id", h.updateBook)
			books.PUT("/:id/authors/:author_id", h.updateBookWithAuthor)
			books.DELETE("/:id", h.deleteBook)
		}

		authors := api.Group("/authors")
		{
			authors.POST("/", h.createAuthor)
			authors.GET("/", h.getAllAuthors)
			authors.GET("/:id", h.getAuthorById)
			authors.PUT("/:id", h.updateAuthor)
			authors.DELETE("/:id", h.deleteAuthor)
		}
	}

	return router
}
