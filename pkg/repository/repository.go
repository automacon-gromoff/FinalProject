package repository

import (
	"database/sql"
	library "github.com/automacon-gromoff/FinalProject"
)

type LibraryAuthor interface {
	Create(author library.Author) (int, error)
	GetAll() ([]library.Author, error)
	GetById(id int) (library.Author, error)
	Update(id int, input library.UpdateAuthorInput) error
	Delete(id int) error
}

type LibraryBook interface {
	Create(book library.Book) (int, error)
	GetAll() ([]library.Book, error)
	GetById(id int) (library.Book, error)
	Update(id int, input library.UpdateBookInput) error
	UpdateWithAuthor(bookId int, authorId int, bookInput library.UpdateBookAndAuthorInput) error
	Delete(id int) error
}

type Repository struct {
	LibraryAuthor
	LibraryBook
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		LibraryAuthor: NewAuthorsPostgres(db),
		LibraryBook:   NewBooksPostgres(db),
	}
}
