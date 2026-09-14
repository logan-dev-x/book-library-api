// Package http
package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/logan-dev-x/book-library-api/internal/book"
)

type Handler struct {
	Service book.Service
}

func (h Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	books := h.Service.GetAll()

	body, err := json.Marshal(books)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(body)
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO no duplicated titles
	var b book.CreateBookInput
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Error: %s", err.Error())
		return
	}

	createdBook, err := h.Service.Create(b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Error: %s", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	body, err := json.Marshal(createdBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Error: %s", err.Error())
		return
	}
	_, _ = w.Write(body)
}
