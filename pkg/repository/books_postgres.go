package repository

import (
	"database/sql"
	"fmt"
	library "github.com/automacon-gromoff/FinalProject"
	"strings"
)

type BooksPostgres struct {
	db *sql.DB
}

func NewBooksPostgres(db *sql.DB) *BooksPostgres {
	return &BooksPostgres{db: db}
}

func (r *BooksPostgres) Create(book library.Book) (int, error) {
	var id int
	createBookQuery := fmt.Sprintf("INSERT INTO %s (name, author_id, publishing_year, isbn) VALUES ($1, $2, $3, $4) RETURNING id", booksTable)
	row := r.db.QueryRow(createBookQuery, book.Name, book.AuthorId, book.PublishingYear, book.ISBN)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *BooksPostgres) GetAll() ([]library.Book, error) {
	var books []library.Book

	getBookQuery := fmt.Sprintf("SELECT id, name, author_id, publishing_year, isbn FROM %s", booksTable)
	rows, err := r.db.Query(getBookQuery)
	defer rows.Close()

	for rows.Next() {
		var book library.Book
		if err := rows.Scan(&book.Id, &book.Name, &book.AuthorId, &book.PublishingYear, &book.ISBN); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, err
}

func (r *BooksPostgres) GetById(id int) (library.Book, error) {
	var book library.Book

	getBookQuery := fmt.Sprintf("SELECT id, name, author_id, publishing_year, isbn FROM %s WHERE id = $1", booksTable)
	err := r.db.QueryRow(getBookQuery, id).Scan(&book.Id, &book.Name, &book.AuthorId, &book.PublishingYear, &book.ISBN)

	return book, err
}

func (r *BooksPostgres) Update(id int, input library.UpdateBookInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.AuthorId != nil {
		setValues = append(setValues, fmt.Sprintf("author_id=$%d", argId))
		args = append(args, *input.AuthorId)
		argId++
	}

	if input.PublishingYear != nil {
		setValues = append(setValues, fmt.Sprintf("publishing_year=$%d", argId))
		args = append(args, *input.PublishingYear)
		argId++
	}

	if input.ISBN != nil {
		setValues = append(setValues, fmt.Sprintf("isbn=$%d", argId))
		args = append(args, *input.ISBN)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", booksTable, setQuery, argId)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *BooksPostgres) UpdateWithAuthor(bookId int, authorId int, input library.UpdateBookAndAuthorInput) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	updateBookQuery := fmt.Sprintf("UPDATE %s SET name=$1 WHERE id = $2", booksTable)
	_, err = r.db.Exec(updateBookQuery, input.BookName, bookId)
	if err != nil {
		tx.Rollback()
		return err
	}

	updateAuthorQuery := fmt.Sprintf("UPDATE %s SET biography=$1 WHERE id = $2", authorsTable)
	_, err = r.db.Exec(updateAuthorQuery, input.AuthorBiography, authorId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *BooksPostgres) Delete(id int) error {
	deleteBookQuery := fmt.Sprintf("DELETE FROM %s WHERE id = $1", booksTable)
	_, err := r.db.Exec(deleteBookQuery, id)

	return err
}
