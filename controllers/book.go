package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"gobooks-api/models"
	"gobooks-api/repository/book"
	"net/http"
	"strconv"
)

type Controller struct{}

var books []models.Book
var repository = bookRepository.BookRepository{}

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r * http.Request) {
		books := repository.GetBooks(db)

		json.NewEncoder(w).Encode(books)
	}
}

func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r * http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		book := repository.GetBook(db, id)

		json.NewEncoder(w).Encode(&book)
	}
}

func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r * http.Request) {
		var book models.Book
		_ = json.NewDecoder(r.Body).Decode(&book)

		bookID := repository.AddBook(db, book)

		json.NewEncoder(w).Encode(bookID)
	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r * http.Request) {
		var book models.Book
		_ = json.NewDecoder(r.Body).Decode(&book)

		rowsUpdated := repository.UpdateBook(db, book)

		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r * http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		rowsDeleted := repository.RemoveBook(db, id)

		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(rowsDeleted)
	}
}
