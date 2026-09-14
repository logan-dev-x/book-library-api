// Package book
package book

import "time"

type Book struct {
	ID int
	Title,
	Author,
	ISBN,
	Description string
	PublishedAt,
	CreatedAt,
	UpdatedAt time.Time
}

type CreateBookInput struct {
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	ISBN        string    `json:"isbn"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"published_at"`
}

type UpdateBookInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
