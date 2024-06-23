package service

import (
	library "github.com/automacon-gromoff/FinalProject"
	"github.com/automacon-gromoff/FinalProject/pkg/repository"
)

type AuthorService struct {
	repo repository.LibraryAuthor
}

func NewAuthorService(repo repository.LibraryAuthor) *AuthorService {
	return &AuthorService{repo: repo}
}

func (s *AuthorService) Create(author library.Author) (int, error) {
	return s.repo.Create(author)
}

func (s *AuthorService) GetAll() ([]library.Author, error) {
	return s.repo.GetAll()
}

func (s *AuthorService) GetById(id int) (library.Author, error) {
	return s.repo.GetById(id)
}

func (s *AuthorService) Update(id int, authorInput library.UpdateAuthorInput) error {
	if err := authorInput.Validate(); err != nil {
		return err
	}
	return s.repo.Update(id, authorInput)
}

func (s *AuthorService) Delete(id int) error {
	return s.repo.Delete(id)
}
