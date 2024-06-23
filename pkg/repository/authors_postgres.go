package repository

import (
	"database/sql"
	"fmt"
	library "github.com/automacon-gromoff/FinalProject"
	"strings"
)

type AuthorsPostgres struct {
	db *sql.DB
}

func NewAuthorsPostgres(db *sql.DB) *AuthorsPostgres {
	return &AuthorsPostgres{db: db}
}

func (r *AuthorsPostgres) Create(author library.Author) (int, error) {
	var id int
	createAuthorQuery := fmt.Sprintf("INSERT INTO %s (name, surname, biography, born_date) VALUES ($1, $2, $3, NULLIF($4,'')::date) RETURNING id", authorsTable)
	row := r.db.QueryRow(createAuthorQuery, author.Name, author.Surname, author.Biography, author.BornDate)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthorsPostgres) GetAll() ([]library.Author, error) {
	var authors []library.Author

	getAllAuthorsQuery := fmt.Sprintf("SELECT id, name, surname, biography, born_date FROM %s", authorsTable)
	rows, err := r.db.Query(getAllAuthorsQuery)
	defer rows.Close()

	for rows.Next() {
		var author library.Author
		rows.Scan(&author.Id, &author.Name, &author.Surname, &author.Biography, &author.BornDate)
		authors = append(authors, author)
	}
	return authors, err
}

func (r *AuthorsPostgres) GetById(id int) (library.Author, error) {
	var author library.Author

	getAuthorQuery := fmt.Sprintf("SELECT id, name, surname, biography, born_date FROM %s WHERE id = $1", authorsTable)
	err := r.db.QueryRow(getAuthorQuery, id).Scan(&author.Id, &author.Name, &author.Surname, &author.Biography, &author.BornDate)

	return author, err
}

func (r *AuthorsPostgres) Update(id int, input library.UpdateAuthorInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Surname != nil {
		setValues = append(setValues, fmt.Sprintf("surname=$%d", argId))
		args = append(args, *input.Surname)
		argId++
	}

	if input.Biography != nil {
		setValues = append(setValues, fmt.Sprintf("biography=$%d", argId))
		args = append(args, *input.Biography)
		argId++
	}

	if input.BornDate != nil {
		setValues = append(setValues, fmt.Sprintf("born_date=$%d", argId))
		args = append(args, *input.BornDate)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", authorsTable, setQuery, argId)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *AuthorsPostgres) Delete(id int) error {
	deleteAuthorQuery := fmt.Sprintf("DELETE FROM %s WHERE id = $1", authorsTable)
	_, err := r.db.Exec(deleteAuthorQuery, id)

	return err
}
