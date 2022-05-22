package routes

import (
	"github.com/gorilla/mux"
	"github.com/prabahar94/bookstore-golan/pkg/controllers"
)

var bookstore_routes = func(router *mux.Routers) {
	router.HandleFunc("/books/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")

}
