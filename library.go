package library

import "errors"

type Book struct {
	Id             int    `json:"id" db:"id"`
	Name           string `json:"name" db:"name" binding:"required"`
	AuthorId       int    `json:"author_id" db:"author_id"`
	PublishingYear int    `json:"publishing_year" db:"publishing_year"`
	ISBN           string `json:"isbn" db:"isbn" binding:"required"`
}

type Author struct {
	Id        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name" binding:"required"`
	Surname   string `json:"surname" db:"surname"`
	Biography string `json:"biography" db:"biography"`
	BornDate  string `json:"born_date" db:"born_date"`
}

type UpdateBookInput struct {
	Name           *string `json:"name"`
	AuthorId       *int    `json:"author_id"`
	PublishingYear *int    `json:"publishing_year"`
	ISBN           *string `json:"isbn"`
}

type UpdateAuthorInput struct {
	Name      *string `json:"name"`
	Surname   *string `json:"surname"`
	Biography *string `json:"biography"`
	BornDate  *string `json:"born_date"`
}

type UpdateBookAndAuthorInput struct {
	BookName        string `json:"name" binding:"required"`
	AuthorBiography string `json:"biography" binding:"required"`
}

func (i UpdateBookInput) Validate() error {
	if i.Name == nil && i.AuthorId == nil && i.PublishingYear == nil && i.ISBN == nil {
		return errors.New("не указано ни одного поля для обновления")
	}

	return nil
}

func (i UpdateAuthorInput) Validate() error {
	if i.Name == nil && i.Surname == nil && i.Biography == nil && i.BornDate == nil {
		return errors.New("не указано ни одного поля для обновления")
	}

	return nil
}
