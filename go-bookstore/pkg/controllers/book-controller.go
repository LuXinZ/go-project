package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-bookstore/pkg/models"
	"go-bookstore/pkg/utils"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(&newBooks)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBookById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id, _ := strconv.Atoi(param["id"])
	println(id)
	newBook, _ := models.GetBookById(int64(id))
	res, _ := json.Marshal(newBook)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	utils.ParseBody(r, &book)
	param := mux.Vars(r)
	id, _ := strconv.Atoi(param["id"])
	b, db := models.GetBookById(int64(id))
	if book.Name != "" {
		b.Name = book.Name
	}
	if book.Author != "" {
		b.Author = book.Author
	}
	if book.Publication != "" {
		b.Publication = book.Publication
	}
	db.Save(&b)
	res, _ := json.Marshal(b)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id, _ := strconv.Atoi(param["id"])
	book := models.DeleteBook(int64(id))
	res, _ := json.Marshal(&book)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	utils.ParseBody(r, &book)
	b := book.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
