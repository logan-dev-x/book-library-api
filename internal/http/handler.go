// Package http
package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/logan-dev-x/book-library-api/internal/book"
)

type Handler struct {
	Service book.Service
}

func NewHandler(service book.Service) Handler {
	return Handler{service}
}

func (h Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Service.Delete(id)
	if err != nil {
		createHeader(w, http.StatusBadRequest)
		_, _ = w.Write(fmt.Appendf([]byte{}, `{"message": "%s"}`, err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	books := h.Service.GetAll()

	body, err := json.Marshal(books)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	createHeader(w, http.StatusOK)
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

	body, err := json.Marshal(createdBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Error: %s", err.Error())
		return
	}

	createHeader(w, http.StatusCreated)
	_, _ = w.Write(body)
}

func createHeader(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
}
