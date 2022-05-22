package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla.mux"
	"strconv"
	"net/http"
	"github.com/prabahar94/bookstore-golang/pkg/utils"
	"github.com/prabahar94/bookstore-golang/pkg/models"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request){
	newbooks :=models.getBooks()
	res, _ := json.Marshal(newbooks)
	w.Header().set("Content-Type","pkglicaiton/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil{
		fmt.Println("Error while parsing")
	}
	bookdetails, _ := models.getBook(ID)
	res, _ := json.Marshal(bookdetails)
	w.Header().set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.parseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _:=json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil{
		fmt.Println("Error while parsing")
	}
	d, _ := models.DeleteBook(ID)
	res, _ := json.Marshal(d)
	w.Header().Set("Content-Type","pkglication/json")
	w.Write(res)
	w.WriteHeader(http.StatusOK)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook := &models.Book{}
	utils.parseBody(r,updateBook)
	vars  := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.parseInt(bookId, 0 , 0)
	if err != nil {
		fmt.Println("Error While parsing ")
	}
	bookdetails, db :=models.getBook(ID)
	if updateBook.Name != "" {
		bookdetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookdetails.Author = updateBook.Author
	}
	if updateBook.publication != ""{
		bookdetails.publication = updateBook.publication
	}
	db.save(&bookdetails)
	res, _ :=json.Marshal(updateBook)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}