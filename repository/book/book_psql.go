package bookRepository

import (
	"database/sql"
	"gobooks-api/utils"
	"gobooks-api/models"
)

type BookRepository struct{}

func (b BookRepository) GetBooks(db *sql.DB) []models.Book {
	var book models.Book
	books := []models.Book{}

	rows, err := db.Query("select * from books")
	utils.LogFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		utils.LogFatal(err)

		books = append(books, book)
	}

	return books
}

func (b BookRepository) GetBook(db *sql.DB, id int) models.Book {
	var book models.Book

	rows := db.QueryRow("select * from books where id=$1", id)

	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	utils.LogFatal(err)

	return book
}

func (b BookRepository) AddBook(db *sql.DB, book models.Book) int {
	var bookID int

	err := db.QueryRow("insert into books(title, author, year) values($1, $2, $3) RETURNING id;",
		book.Title, book.Author, book.Year).Scan(&bookID)
	utils.LogFatal(err)

	return bookID
}

func (b BookRepository) UpdateBook(db *sql.DB, book models.Book) int64 {
	result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id",
		&book.Title, &book.Author, &book.Year, &book.ID)
	utils.LogFatal(err)

	rowsUpdated, err := result.RowsAffected()
	utils.LogFatal(err)

	return rowsUpdated
}

func (b BookRepository) RemoveBook(db *sql.DB, id int) int64 {
	result, err := db.Exec("delete from books where id = $1", id)
	utils.LogFatal(err)

	rowsDeleted, err := result.RowsAffected()
	utils.LogFatal(err)

	return rowsDeleted
}
