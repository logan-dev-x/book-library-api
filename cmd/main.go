package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/logan-dev-x/book-library-api/internal/book"
	myHttp "github.com/logan-dev-x/book-library-api/internal/http"
	"github.com/logan-dev-x/book-library-api/internal/repository"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "books.db")
	if err != nil {
		log.Fatal(err)
	}

	s := book.NewService(repository.NewSQLRepository(db))
	h := myHttp.Handler{Service: s}

	http.HandleFunc("/api/v1/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			h.Create(w, r)
		}
	})
	log.Fatal(http.ListenAndServe(":7070", nil))
}
