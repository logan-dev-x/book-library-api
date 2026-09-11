// Package book
package book

import "time"

type Book struct {
	ID          int
	Title       string
	Author      string
	ISBN        string
	Description string
	PublishedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CreateBookInput struct {
	Title       string
	Author      string
	ISBN        string
	Description string
	PublishedAt time.Time
}

type UpdateBookInput struct {
	Title       string
	Description string
}
