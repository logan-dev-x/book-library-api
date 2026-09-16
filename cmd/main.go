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
	repo := repository.NewSQLRepository(setupDB())
	service := book.NewService(repo)
	handler := myHttp.NewHandler(service)
	router(handler)
	run()
}

func run() {
	log.Fatal(http.ListenAndServe(":7070", nil))
}

func router(handler myHttp.Handler) {
	http.HandleFunc("/api/v1/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			handler.Create(w, r)
		case "GET":
			handler.GetAll(w, r)
		}
	})

	http.HandleFunc("/api/v1/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "DELETE":
			handler.Delete(w, r)
		}
	})
}

func setupDB() *sql.DB {
	db, err := sql.Open("sqlite3", "books.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
