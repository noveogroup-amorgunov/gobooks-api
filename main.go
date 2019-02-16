package main

import (
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"gobooks-api/controllers"
	"gobooks-api/driver"
	"gobooks-api/utils"
	"net/http"
	"os"
)

func init() {
	gotenv.Load()
}

func main() {
	db := driver.ConnectDB()
	router := mux.NewRouter()
	controller := controllers.Controller{}

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")

	port := os.Getenv("PORT")
	utils.LogFatal(http.ListenAndServe(port, router))
}
