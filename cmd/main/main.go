package main

import (
	"log"
	"net/http"
	"github.com/prabahar94/bookstore-golang/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	ru := &routes
	ru.bookstore_routes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9800", r))

}
