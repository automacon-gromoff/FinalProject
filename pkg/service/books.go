package service

import (
	library "github.com/automacon-gromoff/FinalProject"
	"github.com/automacon-gromoff/FinalProject/pkg/repository"
)

type BooksService struct {
	repo repository.LibraryBook
}

func NewBookService(repo repository.LibraryBook) *BooksService {
	return &BooksService{repo: repo}
}

func (s *BooksService) Create(book library.Book) (int, error) {
	return s.repo.Create(book)
}

func (s *BooksService) GetAll() ([]library.Book, error) {
	return s.repo.GetAll()
}

func (s *BooksService) GetById(id int) (library.Book, error) {
	return s.repo.GetById(id)
}

func (s *BooksService) Update(id int, bookInput library.UpdateBookInput) error {
	if err := bookInput.Validate(); err != nil {
		return err
	}

	return s.repo.Update(id, bookInput)
}

func (s *BooksService) UpdateWithAuthor(bookId int, authorId int, input library.UpdateBookAndAuthorInput) error {
	return s.repo.UpdateWithAuthor(bookId, authorId, input)
}

func (s *BooksService) Delete(id int) error {
	return s.repo.Delete(id)
}
