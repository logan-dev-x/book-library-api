// Package repository
package repository

import (
	"database/sql"

	"github.com/logan-dev-x/book-library-api/internal/book"
)

type SQLRepository struct {
	db *sql.DB
}

func NewSQLRepository(db *sql.DB) book.Repository {
	return SQLRepository{db: db}
}

// Delete implements [book.Repository].
func (s SQLRepository) Delete(id int) error {
	_, err := s.db.Exec("DELETE FROM books WHERE id = ?", id)
	return err
}

// GetAll implements [book.Repository].
func (s SQLRepository) GetAll() []book.Book {
	row, _ := s.db.Query("SELECT * FROM books")
	books := []book.Book{}
	for row.Next() {
		var b book.Book
		_ = row.Scan(
			&b.ID,
			&b.Author,
			&b.Title,
			&b.ISBN,
			&b.Description,
			&b.CreatedAt,
			&b.UpdatedAt,
		)
		books = append(books, b)
	}

	return books
}

// GetByID implements [book.Repository].
func (s SQLRepository) GetByID(id int) (book.Book, error) {
	row, err := s.db.Query("SELECT * FROM books WHERE id = ?", id)
	if err != nil {
		return book.Book{}, err
	}
	var b book.Book
	for row.Next() {
		err := row.Scan(
			&b.ID,
			&b.Author,
			&b.Description,
			&b.Title,
			&b.ISBN,
			&b.CreatedAt,
			&b.UpdatedAt,
		)
		if err != nil {
			return book.Book{}, err
		}
	}
	return b, nil
}

// Save implements [book.Repository].
func (s SQLRepository) Save(b book.CreateBookInput) (book.Book, error) {
	_, err := s.db.Exec(
		`INSERT INTO books
		(title, author, description, isbn, published_at, create_at, update_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)`,
		b.Title, b.Author, b.Description, b.ISBN, b.PublishedAt,
	)
	if err != nil {
		return book.Book{}, err
	}
	return book.Book{}, nil
}

// Update implements [book.Repository].
func (s SQLRepository) Update(b book.UpdateBookInput) error {
	_, err := s.db.Exec(
		"UPDATE FROM books SET title = ?, description = ? WHERE id = ?",
		b.Title, b.Description,
	)
	return err
}
