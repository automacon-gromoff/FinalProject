package service

import (
	library "github.com/automacon-gromoff/FinalProject"
	"github.com/automacon-gromoff/FinalProject/pkg/repository"
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

type Service struct {
	LibraryAuthor
	LibraryBook
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		LibraryAuthor: NewAuthorService(repos.LibraryAuthor),
		LibraryBook:   NewBookService(repos.LibraryBook),
	}
}
